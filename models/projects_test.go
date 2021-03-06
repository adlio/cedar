// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testProjects(t *testing.T) {
	t.Parallel()

	query := Projects(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testProjectsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	project := &Project{}
	if err = randomize.Struct(seed, project, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = project.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = project.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Projects(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProjectsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	project := &Project{}
	if err = randomize.Struct(seed, project, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = project.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Projects(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Projects(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProjectsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	project := &Project{}
	if err = randomize.Struct(seed, project, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = project.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ProjectSlice{project}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Projects(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testProjectsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	project := &Project{}
	if err = randomize.Struct(seed, project, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = project.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := ProjectExists(tx, project.ID)
	if err != nil {
		t.Errorf("Unable to check if Project exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ProjectExistsG to return true, but got false.")
	}
}
func testProjectsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	project := &Project{}
	if err = randomize.Struct(seed, project, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = project.Insert(tx); err != nil {
		t.Error(err)
	}

	projectFound, err := FindProject(tx, project.ID)
	if err != nil {
		t.Error(err)
	}

	if projectFound == nil {
		t.Error("want a record, got nil")
	}
}
func testProjectsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	project := &Project{}
	if err = randomize.Struct(seed, project, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = project.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Projects(tx).Bind(project); err != nil {
		t.Error(err)
	}
}

func testProjectsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	project := &Project{}
	if err = randomize.Struct(seed, project, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = project.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Projects(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testProjectsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	projectOne := &Project{}
	projectTwo := &Project{}
	if err = randomize.Struct(seed, projectOne, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}
	if err = randomize.Struct(seed, projectTwo, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = projectOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = projectTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Projects(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testProjectsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	projectOne := &Project{}
	projectTwo := &Project{}
	if err = randomize.Struct(seed, projectOne, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}
	if err = randomize.Struct(seed, projectTwo, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = projectOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = projectTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Projects(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func projectBeforeInsertHook(e boil.Executor, o *Project) error {
	*o = Project{}
	return nil
}

func projectAfterInsertHook(e boil.Executor, o *Project) error {
	*o = Project{}
	return nil
}

func projectAfterSelectHook(e boil.Executor, o *Project) error {
	*o = Project{}
	return nil
}

func projectBeforeUpdateHook(e boil.Executor, o *Project) error {
	*o = Project{}
	return nil
}

func projectAfterUpdateHook(e boil.Executor, o *Project) error {
	*o = Project{}
	return nil
}

func projectBeforeDeleteHook(e boil.Executor, o *Project) error {
	*o = Project{}
	return nil
}

func projectAfterDeleteHook(e boil.Executor, o *Project) error {
	*o = Project{}
	return nil
}

func projectBeforeUpsertHook(e boil.Executor, o *Project) error {
	*o = Project{}
	return nil
}

func projectAfterUpsertHook(e boil.Executor, o *Project) error {
	*o = Project{}
	return nil
}

func testProjectsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Project{}
	o := &Project{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, projectDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Project object: %s", err)
	}

	AddProjectHook(boil.BeforeInsertHook, projectBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	projectBeforeInsertHooks = []ProjectHook{}

	AddProjectHook(boil.AfterInsertHook, projectAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	projectAfterInsertHooks = []ProjectHook{}

	AddProjectHook(boil.AfterSelectHook, projectAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	projectAfterSelectHooks = []ProjectHook{}

	AddProjectHook(boil.BeforeUpdateHook, projectBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	projectBeforeUpdateHooks = []ProjectHook{}

	AddProjectHook(boil.AfterUpdateHook, projectAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	projectAfterUpdateHooks = []ProjectHook{}

	AddProjectHook(boil.BeforeDeleteHook, projectBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	projectBeforeDeleteHooks = []ProjectHook{}

	AddProjectHook(boil.AfterDeleteHook, projectAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	projectAfterDeleteHooks = []ProjectHook{}

	AddProjectHook(boil.BeforeUpsertHook, projectBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	projectBeforeUpsertHooks = []ProjectHook{}

	AddProjectHook(boil.AfterUpsertHook, projectAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	projectAfterUpsertHooks = []ProjectHook{}
}
func testProjectsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	project := &Project{}
	if err = randomize.Struct(seed, project, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = project.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Projects(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProjectsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	project := &Project{}
	if err = randomize.Struct(seed, project, projectDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = project.Insert(tx, projectColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Projects(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProjectToManyParentProjects(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Project
	var b, c Project

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, projectDBTypes, false, projectColumnsWithDefault...)
	randomize.Struct(seed, &c, projectDBTypes, false, projectColumnsWithDefault...)

	b.ParentID.Valid = true
	c.ParentID.Valid = true
	b.ParentID.Int = a.ID
	c.ParentID.Int = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	project, err := a.ParentProjects(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range project {
		if v.ParentID.Int == b.ParentID.Int {
			bFound = true
		}
		if v.ParentID.Int == c.ParentID.Int {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ProjectSlice{&a}
	if err = a.L.LoadParentProjects(tx, false, (*[]*Project)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ParentProjects); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ParentProjects = nil
	if err = a.L.LoadParentProjects(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ParentProjects); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", project)
	}
}

func testProjectToManyTasks(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Project
	var b, c Task

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, taskDBTypes, false, taskColumnsWithDefault...)
	randomize.Struct(seed, &c, taskDBTypes, false, taskColumnsWithDefault...)

	b.ProjectID = a.ID
	c.ProjectID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	task, err := a.Tasks(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range task {
		if v.ProjectID == b.ProjectID {
			bFound = true
		}
		if v.ProjectID == c.ProjectID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ProjectSlice{&a}
	if err = a.L.LoadTasks(tx, false, (*[]*Project)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Tasks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Tasks = nil
	if err = a.L.LoadTasks(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Tasks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", task)
	}
}

func testProjectToManyAddOpParentProjects(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Project
	var b, c, d, e Project

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Project{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Project{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddParentProjects(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ParentID.Int {
			t.Error("foreign key was wrong value", a.ID, first.ParentID.Int)
		}
		if a.ID != second.ParentID.Int {
			t.Error("foreign key was wrong value", a.ID, second.ParentID.Int)
		}

		if first.R.Parent != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Parent != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ParentProjects[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ParentProjects[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ParentProjects(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testProjectToManySetOpParentProjects(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Project
	var b, c, d, e Project

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Project{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetParentProjects(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ParentProjects(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetParentProjects(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ParentProjects(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.ParentID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.ParentID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.ID != d.ParentID.Int {
		t.Error("foreign key was wrong value", a.ID, d.ParentID.Int)
	}
	if a.ID != e.ParentID.Int {
		t.Error("foreign key was wrong value", a.ID, e.ParentID.Int)
	}

	if b.R.Parent != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Parent != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Parent != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Parent != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.ParentProjects[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.ParentProjects[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testProjectToManyRemoveOpParentProjects(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Project
	var b, c, d, e Project

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Project{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddParentProjects(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ParentProjects(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveParentProjects(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ParentProjects(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.ParentID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.ParentID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Parent != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Parent != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Parent != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Parent != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.ParentProjects) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.ParentProjects[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.ParentProjects[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testProjectToManyAddOpTasks(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Project
	var b, c, d, e Task

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Task{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, taskDBTypes, false, strmangle.SetComplement(taskPrimaryKeyColumns, taskColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Task{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTasks(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ProjectID {
			t.Error("foreign key was wrong value", a.ID, first.ProjectID)
		}
		if a.ID != second.ProjectID {
			t.Error("foreign key was wrong value", a.ID, second.ProjectID)
		}

		if first.R.Project != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Project != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Tasks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Tasks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Tasks(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testProjectToOneTenantUsingTenant(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Project
	var foreign Tenant

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, tenantDBTypes, false, tenantColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tenant struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.TenantID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Tenant(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ProjectSlice{&local}
	if err = local.L.LoadTenant(tx, false, (*[]*Project)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Tenant == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Tenant = nil
	if err = local.L.LoadTenant(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Tenant == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testProjectToOneClientUsingClient(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Project
	var foreign Client

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, clientDBTypes, false, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ClientID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Client(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ProjectSlice{&local}
	if err = local.L.LoadClient(tx, false, (*[]*Project)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Client == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Client = nil
	if err = local.L.LoadClient(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Client == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testProjectToOneProjectUsingParent(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Project
	var foreign Project

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	local.ParentID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ParentID.Int = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Parent(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ProjectSlice{&local}
	if err = local.L.LoadParent(tx, false, (*[]*Project)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Parent == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Parent = nil
	if err = local.L.LoadParent(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Parent == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testProjectToOneSetOpTenantUsingTenant(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Project
	var b, c Tenant

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, tenantDBTypes, false, strmangle.SetComplement(tenantPrimaryKeyColumns, tenantColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, tenantDBTypes, false, strmangle.SetComplement(tenantPrimaryKeyColumns, tenantColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Tenant{&b, &c} {
		err = a.SetTenant(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Tenant != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Projects[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TenantID != x.ID {
			t.Error("foreign key was wrong value", a.TenantID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TenantID))
		reflect.Indirect(reflect.ValueOf(&a.TenantID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TenantID != x.ID {
			t.Error("foreign key was wrong value", a.TenantID, x.ID)
		}
	}
}
func testProjectToOneSetOpClientUsingClient(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Project
	var b, c Client

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, clientDBTypes, false, strmangle.SetComplement(clientPrimaryKeyColumns, clientColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, clientDBTypes, false, strmangle.SetComplement(clientPrimaryKeyColumns, clientColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Client{&b, &c} {
		err = a.SetClient(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Client != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Projects[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ClientID != x.ID {
			t.Error("foreign key was wrong value", a.ClientID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ClientID))
		reflect.Indirect(reflect.ValueOf(&a.ClientID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ClientID != x.ID {
			t.Error("foreign key was wrong value", a.ClientID, x.ID)
		}
	}
}
func testProjectToOneSetOpProjectUsingParent(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Project
	var b, c Project

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Project{&b, &c} {
		err = a.SetParent(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Parent != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ParentProjects[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ParentID.Int != x.ID {
			t.Error("foreign key was wrong value", a.ParentID.Int)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ParentID.Int))
		reflect.Indirect(reflect.ValueOf(&a.ParentID.Int)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ParentID.Int != x.ID {
			t.Error("foreign key was wrong value", a.ParentID.Int, x.ID)
		}
	}
}

func testProjectToOneRemoveOpProjectUsingParent(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Project
	var b Project

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetParent(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveParent(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Parent(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Parent != nil {
		t.Error("R struct entry should be nil")
	}

	if a.ParentID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.ParentProjects) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testProjectsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	project := &Project{}
	if err = randomize.Struct(seed, project, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = project.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = project.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testProjectsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	project := &Project{}
	if err = randomize.Struct(seed, project, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = project.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ProjectSlice{project}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testProjectsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	project := &Project{}
	if err = randomize.Struct(seed, project, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = project.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Projects(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	projectDBTypes = map[string]string{`ClientID`: `integer`, `CreatedAt`: `timestamp with time zone`, `EndsOn`: `date`, `ID`: `integer`, `IsActive`: `boolean`, `IsBillable`: `boolean`, `Name`: `text`, `ParentID`: `integer`, `Slug`: `text`, `StartsOn`: `date`, `TenantID`: `integer`, `UpdatedAt`: `timestamp with time zone`}
	_              = bytes.MinRead
)

func testProjectsUpdate(t *testing.T) {
	t.Parallel()

	if len(projectColumns) == len(projectPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	project := &Project{}
	if err = randomize.Struct(seed, project, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = project.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Projects(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, project, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	if err = project.Update(tx); err != nil {
		t.Error(err)
	}
}

func testProjectsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(projectColumns) == len(projectPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	project := &Project{}
	if err = randomize.Struct(seed, project, projectDBTypes, true, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = project.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Projects(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, project, projectDBTypes, true, projectPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(projectColumns, projectPrimaryKeyColumns) {
		fields = projectColumns
	} else {
		fields = strmangle.SetComplement(
			projectColumns,
			projectPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(project))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := ProjectSlice{project}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testProjectsUpsert(t *testing.T) {
	t.Parallel()

	if len(projectColumns) == len(projectPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	project := Project{}
	if err = randomize.Struct(seed, &project, projectDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = project.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Project: %s", err)
	}

	count, err := Projects(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &project, projectDBTypes, false, projectPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	if err = project.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Project: %s", err)
	}

	count, err = Projects(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
