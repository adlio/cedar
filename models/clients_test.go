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

func testClients(t *testing.T) {
	t.Parallel()

	query := Clients(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testClientsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	client := &Client{}
	if err = randomize.Struct(seed, client, clientDBTypes, true, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = client.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = client.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Clients(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClientsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	client := &Client{}
	if err = randomize.Struct(seed, client, clientDBTypes, true, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = client.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Clients(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Clients(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClientsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	client := &Client{}
	if err = randomize.Struct(seed, client, clientDBTypes, true, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = client.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ClientSlice{client}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Clients(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testClientsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	client := &Client{}
	if err = randomize.Struct(seed, client, clientDBTypes, true, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = client.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := ClientExists(tx, client.ID)
	if err != nil {
		t.Errorf("Unable to check if Client exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ClientExistsG to return true, but got false.")
	}
}
func testClientsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	client := &Client{}
	if err = randomize.Struct(seed, client, clientDBTypes, true, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = client.Insert(tx); err != nil {
		t.Error(err)
	}

	clientFound, err := FindClient(tx, client.ID)
	if err != nil {
		t.Error(err)
	}

	if clientFound == nil {
		t.Error("want a record, got nil")
	}
}
func testClientsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	client := &Client{}
	if err = randomize.Struct(seed, client, clientDBTypes, true, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = client.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Clients(tx).Bind(client); err != nil {
		t.Error(err)
	}
}

func testClientsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	client := &Client{}
	if err = randomize.Struct(seed, client, clientDBTypes, true, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = client.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Clients(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testClientsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	clientOne := &Client{}
	clientTwo := &Client{}
	if err = randomize.Struct(seed, clientOne, clientDBTypes, false, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}
	if err = randomize.Struct(seed, clientTwo, clientDBTypes, false, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = clientOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = clientTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Clients(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testClientsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	clientOne := &Client{}
	clientTwo := &Client{}
	if err = randomize.Struct(seed, clientOne, clientDBTypes, false, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}
	if err = randomize.Struct(seed, clientTwo, clientDBTypes, false, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = clientOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = clientTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Clients(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func clientBeforeInsertHook(e boil.Executor, o *Client) error {
	*o = Client{}
	return nil
}

func clientAfterInsertHook(e boil.Executor, o *Client) error {
	*o = Client{}
	return nil
}

func clientAfterSelectHook(e boil.Executor, o *Client) error {
	*o = Client{}
	return nil
}

func clientBeforeUpdateHook(e boil.Executor, o *Client) error {
	*o = Client{}
	return nil
}

func clientAfterUpdateHook(e boil.Executor, o *Client) error {
	*o = Client{}
	return nil
}

func clientBeforeDeleteHook(e boil.Executor, o *Client) error {
	*o = Client{}
	return nil
}

func clientAfterDeleteHook(e boil.Executor, o *Client) error {
	*o = Client{}
	return nil
}

func clientBeforeUpsertHook(e boil.Executor, o *Client) error {
	*o = Client{}
	return nil
}

func clientAfterUpsertHook(e boil.Executor, o *Client) error {
	*o = Client{}
	return nil
}

func testClientsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Client{}
	o := &Client{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, clientDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Client object: %s", err)
	}

	AddClientHook(boil.BeforeInsertHook, clientBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	clientBeforeInsertHooks = []ClientHook{}

	AddClientHook(boil.AfterInsertHook, clientAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	clientAfterInsertHooks = []ClientHook{}

	AddClientHook(boil.AfterSelectHook, clientAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	clientAfterSelectHooks = []ClientHook{}

	AddClientHook(boil.BeforeUpdateHook, clientBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	clientBeforeUpdateHooks = []ClientHook{}

	AddClientHook(boil.AfterUpdateHook, clientAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	clientAfterUpdateHooks = []ClientHook{}

	AddClientHook(boil.BeforeDeleteHook, clientBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	clientBeforeDeleteHooks = []ClientHook{}

	AddClientHook(boil.AfterDeleteHook, clientAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	clientAfterDeleteHooks = []ClientHook{}

	AddClientHook(boil.BeforeUpsertHook, clientBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	clientBeforeUpsertHooks = []ClientHook{}

	AddClientHook(boil.AfterUpsertHook, clientAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	clientAfterUpsertHooks = []ClientHook{}
}
func testClientsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	client := &Client{}
	if err = randomize.Struct(seed, client, clientDBTypes, true, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = client.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Clients(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClientsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	client := &Client{}
	if err = randomize.Struct(seed, client, clientDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = client.Insert(tx, clientColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Clients(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClientToManyParentClients(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Client
	var b, c Client

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, clientDBTypes, true, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, clientDBTypes, false, clientColumnsWithDefault...)
	randomize.Struct(seed, &c, clientDBTypes, false, clientColumnsWithDefault...)

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

	client, err := a.ParentClients(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range client {
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

	slice := ClientSlice{&a}
	if err = a.L.LoadParentClients(tx, false, (*[]*Client)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ParentClients); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ParentClients = nil
	if err = a.L.LoadParentClients(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ParentClients); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", client)
	}
}

func testClientToManyProjects(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Client
	var b, c Project

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, clientDBTypes, true, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, projectDBTypes, false, projectColumnsWithDefault...)
	randomize.Struct(seed, &c, projectDBTypes, false, projectColumnsWithDefault...)

	b.ClientID = a.ID
	c.ClientID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	project, err := a.Projects(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range project {
		if v.ClientID == b.ClientID {
			bFound = true
		}
		if v.ClientID == c.ClientID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ClientSlice{&a}
	if err = a.L.LoadProjects(tx, false, (*[]*Client)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Projects); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Projects = nil
	if err = a.L.LoadProjects(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Projects); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", project)
	}
}

func testClientToManyAddOpParentClients(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Client
	var b, c, d, e Client

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, clientDBTypes, false, strmangle.SetComplement(clientPrimaryKeyColumns, clientColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Client{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, clientDBTypes, false, strmangle.SetComplement(clientPrimaryKeyColumns, clientColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Client{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddParentClients(tx, i != 0, x...)
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

		if a.R.ParentClients[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ParentClients[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ParentClients(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testClientToManySetOpParentClients(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Client
	var b, c, d, e Client

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, clientDBTypes, false, strmangle.SetComplement(clientPrimaryKeyColumns, clientColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Client{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, clientDBTypes, false, strmangle.SetComplement(clientPrimaryKeyColumns, clientColumnsWithoutDefault)...); err != nil {
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

	err = a.SetParentClients(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ParentClients(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetParentClients(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ParentClients(tx).Count()
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

	if a.R.ParentClients[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.ParentClients[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testClientToManyRemoveOpParentClients(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Client
	var b, c, d, e Client

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, clientDBTypes, false, strmangle.SetComplement(clientPrimaryKeyColumns, clientColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Client{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, clientDBTypes, false, strmangle.SetComplement(clientPrimaryKeyColumns, clientColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddParentClients(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ParentClients(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveParentClients(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ParentClients(tx).Count()
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

	if len(a.R.ParentClients) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.ParentClients[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.ParentClients[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testClientToManyAddOpProjects(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Client
	var b, c, d, e Project

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, clientDBTypes, false, strmangle.SetComplement(clientPrimaryKeyColumns, clientColumnsWithoutDefault)...); err != nil {
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
		err = a.AddProjects(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ClientID {
			t.Error("foreign key was wrong value", a.ID, first.ClientID)
		}
		if a.ID != second.ClientID {
			t.Error("foreign key was wrong value", a.ID, second.ClientID)
		}

		if first.R.Client != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Client != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Projects[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Projects[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Projects(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testClientToOneTenantUsingTenant(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Client
	var foreign Tenant

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, clientDBTypes, false, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
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

	slice := ClientSlice{&local}
	if err = local.L.LoadTenant(tx, false, (*[]*Client)(&slice)); err != nil {
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

func testClientToOneClientUsingParent(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Client
	var foreign Client

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, clientDBTypes, true, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, clientDBTypes, false, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
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

	slice := ClientSlice{&local}
	if err = local.L.LoadParent(tx, false, (*[]*Client)(&slice)); err != nil {
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

func testClientToOneSetOpTenantUsingTenant(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Client
	var b, c Tenant

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, clientDBTypes, false, strmangle.SetComplement(clientPrimaryKeyColumns, clientColumnsWithoutDefault)...); err != nil {
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

		if x.R.Clients[0] != &a {
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
func testClientToOneSetOpClientUsingParent(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Client
	var b, c Client

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, clientDBTypes, false, strmangle.SetComplement(clientPrimaryKeyColumns, clientColumnsWithoutDefault)...); err != nil {
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
		err = a.SetParent(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Parent != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ParentClients[0] != &a {
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

func testClientToOneRemoveOpClientUsingParent(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Client
	var b Client

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, clientDBTypes, false, strmangle.SetComplement(clientPrimaryKeyColumns, clientColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, clientDBTypes, false, strmangle.SetComplement(clientPrimaryKeyColumns, clientColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ParentClients) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testClientsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	client := &Client{}
	if err = randomize.Struct(seed, client, clientDBTypes, true, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = client.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = client.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testClientsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	client := &Client{}
	if err = randomize.Struct(seed, client, clientDBTypes, true, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = client.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ClientSlice{client}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testClientsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	client := &Client{}
	if err = randomize.Struct(seed, client, clientDBTypes, true, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = client.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Clients(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	clientDBTypes = map[string]string{`CreatedAt`: `timestamp with time zone`, `ID`: `integer`, `IsActive`: `boolean`, `IsBillable`: `boolean`, `Name`: `text`, `ParentID`: `integer`, `Slug`: `text`, `TenantID`: `integer`, `UpdatedAt`: `timestamp with time zone`}
	_             = bytes.MinRead
)

func testClientsUpdate(t *testing.T) {
	t.Parallel()

	if len(clientColumns) == len(clientPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	client := &Client{}
	if err = randomize.Struct(seed, client, clientDBTypes, true, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = client.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Clients(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, client, clientDBTypes, true, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	if err = client.Update(tx); err != nil {
		t.Error(err)
	}
}

func testClientsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(clientColumns) == len(clientPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	client := &Client{}
	if err = randomize.Struct(seed, client, clientDBTypes, true, clientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = client.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Clients(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, client, clientDBTypes, true, clientPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(clientColumns, clientPrimaryKeyColumns) {
		fields = clientColumns
	} else {
		fields = strmangle.SetComplement(
			clientColumns,
			clientPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(client))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := ClientSlice{client}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testClientsUpsert(t *testing.T) {
	t.Parallel()

	if len(clientColumns) == len(clientPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	client := Client{}
	if err = randomize.Struct(seed, &client, clientDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = client.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Client: %s", err)
	}

	count, err := Clients(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &client, clientDBTypes, false, clientPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Client struct: %s", err)
	}

	if err = client.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Client: %s", err)
	}

	count, err = Clients(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}