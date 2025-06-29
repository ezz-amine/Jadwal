package output

import (
	"context"
	"fmt"
	"strings"

	"github.com/ezz-amine/Jadwal/pkg/core"
	"github.com/ezz-amine/Jadwal/pkg/sqlc"
)

const (
	TODO_SEPARATOR           string = "|"
	TODO_MARKER__DONE        string = "[]"
	TODO_MARKER__NOT_DONE    string = "[ ]"
	TODO_MARKER__PLACEHOLDER string = "   "
	TITLE_PADDING_CHAR       string = "="
	FIXED_COLUMN_SIZE        int    = 4
)

func displayTableTitle(termContext *core.TermContext, table sqlc.TodoTable) {
	title := fmt.Sprintf(" %s ", table.Title)
	title = core.CenterPad(title, '=', termContext.UsedWidth)

	fmt.Println(title)
}

func displayEntryAsLine(termContext *core.TermContext, entry sqlc.TodoEntry, entryPosition int) {
	marker := TODO_MARKER__NOT_DONE

	if entry.IsDone.Bool && entry.IsDone.Valid {
		marker = TODO_MARKER__DONE
	}

	textWidth := termContext.UsedWidth - FIXED_COLUMN_SIZE*2 - 2 // the full width - double the fixed column size (4*2)  - 2 (spaces around the text)
	splitedContent := core.SplitTextByWords(entry.Content, textWidth)
	firstLineFormat := "%[1]s%[2]s %[3]s %[2]s%02[4]d%[2]s\n"
	otherLineFormat := "%3[1]s%[2]s %[3]s %[2]s%2[1]s%[2]s\n"

	for i, line := range splitedContent {
		line = fmt.Sprintf("%-*s", textWidth, line)

		if i == 0 {
			fmt.Printf(firstLineFormat, marker, TODO_SEPARATOR, line, entryPosition)
			continue
		}

		fmt.Printf(otherLineFormat, " ", TODO_SEPARATOR, line)
	}
}

func DisplayAllTables(termContext *core.TermContext, tablesLister core.TablesLister) error {
	ctx, cancel := context.WithTimeout(context.Background(), core.TIMEOUT)
	defer cancel()

	tables, err := tablesLister.ListTables(ctx)
	if err != nil {
		return fmt.Errorf("can't display tables: %w", err)
	}

	for idx, table := range tables {
		if idx > 0 {
			fmt.Println("")
		}
		err = DisplayTable(termContext, tablesLister, table)
		if err != nil {
			return err
		}
	}

	return nil
}

func DisplayTable(termContext *core.TermContext, entriesLister core.EntriesLister, table sqlc.TodoTable) error {
	ctx, cancel := context.WithTimeout(context.Background(), core.TIMEOUT)
	defer cancel()

	displayTableTitle(termContext, table)

	todos, err := entriesLister.ListEntries(ctx, table.ID)
	if err != nil {
		return fmt.Errorf("can't get todos from '%s': %w", table.Title, err)
	}

	if len(todos) > 0 {
		for idx, todo := range todos {
			displayEntryAsLine(termContext, todo, idx)
		}
	} else {
		fmt.Printf(" - %-*s\n", termContext.UsedWidth-3, "empty table: no TODO to show")
	}

	fmt.Println(strings.Repeat("=", termContext.UsedWidth))
	return nil
}
