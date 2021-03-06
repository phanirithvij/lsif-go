package indexer

import (
	"testing"
)

func TestFindDocstringFunc(t *testing.T) {
	packages := getTestPackages(t)
	p, obj := findDefinitionByName(t, packages, "ParallelizableFunc")

	expectedText := normalizeDocstring(`
		ParallelizableFunc is a function that can be called concurrently with other instances
		of this function type.
	`)
	if text := normalizeDocstring(findDocstring(NewPackageDataCache(), packages, p, obj)); text != expectedText {
		t.Errorf("unexpected hover text. want=%q have=%q", expectedText, text)
	}
}

func TestFindDocstringInterface(t *testing.T) {
	packages := getTestPackages(t)
	p, obj := findDefinitionByName(t, packages, "TestInterface")

	expectedText := normalizeDocstring(`TestInterface is an interface used for testing.`)
	if text := normalizeDocstring(findDocstring(NewPackageDataCache(), packages, p, obj)); text != expectedText {
		t.Errorf("unexpected hover text. want=%q have=%q", expectedText, text)
	}
}

func TestFindDocstringStruct(t *testing.T) {
	packages := getTestPackages(t)
	p, obj := findDefinitionByName(t, packages, "TestStruct")

	expectedText := normalizeDocstring(`TestStruct is a struct used for testing.`)
	if text := normalizeDocstring(findDocstring(NewPackageDataCache(), packages, p, obj)); text != expectedText {
		t.Errorf("unexpected hover text. want=%q have=%q", expectedText, text)
	}
}

func TestFindDocstringField(t *testing.T) {
	packages := getTestPackages(t)
	p, obj := findDefinitionByName(t, packages, "NestedC")

	expectedText := normalizeDocstring(`NestedC docs`)
	if text := normalizeDocstring(findDocstring(NewPackageDataCache(), packages, p, obj)); text != expectedText {
		t.Errorf("unexpected hover text. want=%q have=%q", expectedText, text)
	}
}

func TestFindDocstringConst(t *testing.T) {
	packages := getTestPackages(t)
	p, obj := findDefinitionByName(t, packages, "Score")

	expectedText := normalizeDocstring(`Score is just a hardcoded number.`)
	if text := normalizeDocstring(findDocstring(NewPackageDataCache(), packages, p, obj)); text != expectedText {
		t.Errorf("unexpected hover text. want=%q have=%q", expectedText, text)
	}
}

// TestFindDocstringLocalVariable ensures that local definitions within a function with a
// docstring do not take their parent's docstring as their own. This was a brief (unpublished)
// regression made when switching from storing node paths for hover text extraction to only
// storing a single ancestor node from which hover text is extracted.
func TestFindDocstringLocalVariable(t *testing.T) {
	packages := getTestPackages(t)
	p, obj := findDefinitionByName(t, packages, "errs")

	expectedText := normalizeDocstring(``)
	if text := normalizeDocstring(findDocstring(NewPackageDataCache(), packages, p, obj)); text != expectedText {
		t.Errorf("unexpected hover text. want=%q have=%q", expectedText, text)
	}
}

func TestFindDocstringInternalPackageName(t *testing.T) {
	packages := getTestPackages(t)
	p, obj := findUseByName(t, packages, "secret")

	expectedText := normalizeDocstring(`secret is a package that holds secrets.`)
	if text := normalizeDocstring(findDocstring(NewPackageDataCache(), packages, p, obj)); text != expectedText {
		t.Errorf("unexpected hover text. want=%q have=%q", expectedText, text)
	}
}

func TestFindDocstringExternalPackageName(t *testing.T) {
	packages := getTestPackages(t)
	p, obj := findUseByName(t, packages, "sync")

	expectedText := normalizeDocstring(`
		Package sync provides basic synchronization primitives such as mutual exclusion locks.
		Other than the Once and WaitGroup types, most are intended for use by low-level library routines.
		Higher-level synchronization is better done via channels and communication.
		Values containing the types defined in this package should not be copied.
	`)
	if text := normalizeDocstring(findDocstring(NewPackageDataCache(), packages, p, obj)); text != expectedText {
		t.Errorf("unexpected hover text. want=%q have=%q", expectedText, text)
	}
}

func TestFindExternalDocstring(t *testing.T) {
	packages := getTestPackages(t)
	p, obj := findUseByName(t, packages, "WaitGroup")

	expectedText := normalizeDocstring(`
		A WaitGroup waits for a collection of goroutines to finish.
		The main goroutine calls Add to set the number of goroutines to wait for.
		Then each of the goroutines runs and calls Done when finished.
		At the same time, Wait can be used to block until all goroutines have finished.
		A WaitGroup must not be copied after first use.
	`)
	if text := normalizeDocstring(findExternalDocstring(NewPackageDataCache(), packages, p, obj)); text != expectedText {
		t.Errorf("unexpected hover text. want=%q have=%q", expectedText, text)
	}
}
