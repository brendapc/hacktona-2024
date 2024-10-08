// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ShelterNeed is an object representing the database table.
type ShelterNeed struct {
	ID                int  `boil:"id" json:"id" toml:"id" yaml:"id"`
	ShelterID         int  `boil:"shelter_id" json:"shelter_id" toml:"shelter_id" yaml:"shelter_id"`
	BeddingItem       bool `boil:"bedding_item" json:"bedding_item" toml:"bedding_item" yaml:"bedding_item"`
	FoodNonPerishable bool `boil:"food_non_perishable" json:"food_non_perishable" toml:"food_non_perishable" yaml:"food_non_perishable"`
	FoodPerishable    bool `boil:"food_perishable" json:"food_perishable" toml:"food_perishable" yaml:"food_perishable"`
	HygieneProducts   bool `boil:"hygiene_products" json:"hygiene_products" toml:"hygiene_products" yaml:"hygiene_products"`
	ClothingMale      bool `boil:"clothing_male" json:"clothing_male" toml:"clothing_male" yaml:"clothing_male"`
	ClothingFemale    bool `boil:"clothing_female" json:"clothing_female" toml:"clothing_female" yaml:"clothing_female"`
	ClothingChildren  bool `boil:"clothing_children" json:"clothing_children" toml:"clothing_children" yaml:"clothing_children"`
	MedicalSupplies   bool `boil:"medical_supplies" json:"medical_supplies" toml:"medical_supplies" yaml:"medical_supplies"`
	PetFoodDogs       bool `boil:"pet_food_dogs" json:"pet_food_dogs" toml:"pet_food_dogs" yaml:"pet_food_dogs"`
	PetFoodCats       bool `boil:"pet_food_cats" json:"pet_food_cats" toml:"pet_food_cats" yaml:"pet_food_cats"`
	CleaningSupplies  bool `boil:"cleaning_supplies" json:"cleaning_supplies" toml:"cleaning_supplies" yaml:"cleaning_supplies"`

	R *shelterNeedR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L shelterNeedL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ShelterNeedColumns = struct {
	ID                string
	ShelterID         string
	BeddingItem       string
	FoodNonPerishable string
	FoodPerishable    string
	HygieneProducts   string
	ClothingMale      string
	ClothingFemale    string
	ClothingChildren  string
	MedicalSupplies   string
	PetFoodDogs       string
	PetFoodCats       string
	CleaningSupplies  string
}{
	ID:                "id",
	ShelterID:         "shelter_id",
	BeddingItem:       "bedding_item",
	FoodNonPerishable: "food_non_perishable",
	FoodPerishable:    "food_perishable",
	HygieneProducts:   "hygiene_products",
	ClothingMale:      "clothing_male",
	ClothingFemale:    "clothing_female",
	ClothingChildren:  "clothing_children",
	MedicalSupplies:   "medical_supplies",
	PetFoodDogs:       "pet_food_dogs",
	PetFoodCats:       "pet_food_cats",
	CleaningSupplies:  "cleaning_supplies",
}

var ShelterNeedTableColumns = struct {
	ID                string
	ShelterID         string
	BeddingItem       string
	FoodNonPerishable string
	FoodPerishable    string
	HygieneProducts   string
	ClothingMale      string
	ClothingFemale    string
	ClothingChildren  string
	MedicalSupplies   string
	PetFoodDogs       string
	PetFoodCats       string
	CleaningSupplies  string
}{
	ID:                "shelter_needs.id",
	ShelterID:         "shelter_needs.shelter_id",
	BeddingItem:       "shelter_needs.bedding_item",
	FoodNonPerishable: "shelter_needs.food_non_perishable",
	FoodPerishable:    "shelter_needs.food_perishable",
	HygieneProducts:   "shelter_needs.hygiene_products",
	ClothingMale:      "shelter_needs.clothing_male",
	ClothingFemale:    "shelter_needs.clothing_female",
	ClothingChildren:  "shelter_needs.clothing_children",
	MedicalSupplies:   "shelter_needs.medical_supplies",
	PetFoodDogs:       "shelter_needs.pet_food_dogs",
	PetFoodCats:       "shelter_needs.pet_food_cats",
	CleaningSupplies:  "shelter_needs.cleaning_supplies",
}

// Generated where

var ShelterNeedWhere = struct {
	ID                whereHelperint
	ShelterID         whereHelperint
	BeddingItem       whereHelperbool
	FoodNonPerishable whereHelperbool
	FoodPerishable    whereHelperbool
	HygieneProducts   whereHelperbool
	ClothingMale      whereHelperbool
	ClothingFemale    whereHelperbool
	ClothingChildren  whereHelperbool
	MedicalSupplies   whereHelperbool
	PetFoodDogs       whereHelperbool
	PetFoodCats       whereHelperbool
	CleaningSupplies  whereHelperbool
}{
	ID:                whereHelperint{field: "\"shelter_needs\".\"id\""},
	ShelterID:         whereHelperint{field: "\"shelter_needs\".\"shelter_id\""},
	BeddingItem:       whereHelperbool{field: "\"shelter_needs\".\"bedding_item\""},
	FoodNonPerishable: whereHelperbool{field: "\"shelter_needs\".\"food_non_perishable\""},
	FoodPerishable:    whereHelperbool{field: "\"shelter_needs\".\"food_perishable\""},
	HygieneProducts:   whereHelperbool{field: "\"shelter_needs\".\"hygiene_products\""},
	ClothingMale:      whereHelperbool{field: "\"shelter_needs\".\"clothing_male\""},
	ClothingFemale:    whereHelperbool{field: "\"shelter_needs\".\"clothing_female\""},
	ClothingChildren:  whereHelperbool{field: "\"shelter_needs\".\"clothing_children\""},
	MedicalSupplies:   whereHelperbool{field: "\"shelter_needs\".\"medical_supplies\""},
	PetFoodDogs:       whereHelperbool{field: "\"shelter_needs\".\"pet_food_dogs\""},
	PetFoodCats:       whereHelperbool{field: "\"shelter_needs\".\"pet_food_cats\""},
	CleaningSupplies:  whereHelperbool{field: "\"shelter_needs\".\"cleaning_supplies\""},
}

// ShelterNeedRels is where relationship names are stored.
var ShelterNeedRels = struct {
	Shelter string
}{
	Shelter: "Shelter",
}

// shelterNeedR is where relationships are stored.
type shelterNeedR struct {
	Shelter *Shelter `boil:"Shelter" json:"Shelter" toml:"Shelter" yaml:"Shelter"`
}

// NewStruct creates a new relationship struct
func (*shelterNeedR) NewStruct() *shelterNeedR {
	return &shelterNeedR{}
}

func (r *shelterNeedR) GetShelter() *Shelter {
	if r == nil {
		return nil
	}
	return r.Shelter
}

// shelterNeedL is where Load methods for each relationship are stored.
type shelterNeedL struct{}

var (
	shelterNeedAllColumns            = []string{"id", "shelter_id", "bedding_item", "food_non_perishable", "food_perishable", "hygiene_products", "clothing_male", "clothing_female", "clothing_children", "medical_supplies", "pet_food_dogs", "pet_food_cats", "cleaning_supplies"}
	shelterNeedColumnsWithoutDefault = []string{"shelter_id"}
	shelterNeedColumnsWithDefault    = []string{"id", "bedding_item", "food_non_perishable", "food_perishable", "hygiene_products", "clothing_male", "clothing_female", "clothing_children", "medical_supplies", "pet_food_dogs", "pet_food_cats", "cleaning_supplies"}
	shelterNeedPrimaryKeyColumns     = []string{"id"}
	shelterNeedGeneratedColumns      = []string{}
)

type (
	// ShelterNeedSlice is an alias for a slice of pointers to ShelterNeed.
	// This should almost always be used instead of []ShelterNeed.
	ShelterNeedSlice []*ShelterNeed
	// ShelterNeedHook is the signature for custom ShelterNeed hook methods
	ShelterNeedHook func(context.Context, boil.ContextExecutor, *ShelterNeed) error

	shelterNeedQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	shelterNeedType                 = reflect.TypeOf(&ShelterNeed{})
	shelterNeedMapping              = queries.MakeStructMapping(shelterNeedType)
	shelterNeedPrimaryKeyMapping, _ = queries.BindMapping(shelterNeedType, shelterNeedMapping, shelterNeedPrimaryKeyColumns)
	shelterNeedInsertCacheMut       sync.RWMutex
	shelterNeedInsertCache          = make(map[string]insertCache)
	shelterNeedUpdateCacheMut       sync.RWMutex
	shelterNeedUpdateCache          = make(map[string]updateCache)
	shelterNeedUpsertCacheMut       sync.RWMutex
	shelterNeedUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var shelterNeedAfterSelectMu sync.Mutex
var shelterNeedAfterSelectHooks []ShelterNeedHook

var shelterNeedBeforeInsertMu sync.Mutex
var shelterNeedBeforeInsertHooks []ShelterNeedHook
var shelterNeedAfterInsertMu sync.Mutex
var shelterNeedAfterInsertHooks []ShelterNeedHook

var shelterNeedBeforeUpdateMu sync.Mutex
var shelterNeedBeforeUpdateHooks []ShelterNeedHook
var shelterNeedAfterUpdateMu sync.Mutex
var shelterNeedAfterUpdateHooks []ShelterNeedHook

var shelterNeedBeforeDeleteMu sync.Mutex
var shelterNeedBeforeDeleteHooks []ShelterNeedHook
var shelterNeedAfterDeleteMu sync.Mutex
var shelterNeedAfterDeleteHooks []ShelterNeedHook

var shelterNeedBeforeUpsertMu sync.Mutex
var shelterNeedBeforeUpsertHooks []ShelterNeedHook
var shelterNeedAfterUpsertMu sync.Mutex
var shelterNeedAfterUpsertHooks []ShelterNeedHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ShelterNeed) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shelterNeedAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ShelterNeed) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shelterNeedBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ShelterNeed) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shelterNeedAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ShelterNeed) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shelterNeedBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ShelterNeed) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shelterNeedAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ShelterNeed) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shelterNeedBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ShelterNeed) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shelterNeedAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ShelterNeed) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shelterNeedBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ShelterNeed) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shelterNeedAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddShelterNeedHook registers your hook function for all future operations.
func AddShelterNeedHook(hookPoint boil.HookPoint, shelterNeedHook ShelterNeedHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		shelterNeedAfterSelectMu.Lock()
		shelterNeedAfterSelectHooks = append(shelterNeedAfterSelectHooks, shelterNeedHook)
		shelterNeedAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		shelterNeedBeforeInsertMu.Lock()
		shelterNeedBeforeInsertHooks = append(shelterNeedBeforeInsertHooks, shelterNeedHook)
		shelterNeedBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		shelterNeedAfterInsertMu.Lock()
		shelterNeedAfterInsertHooks = append(shelterNeedAfterInsertHooks, shelterNeedHook)
		shelterNeedAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		shelterNeedBeforeUpdateMu.Lock()
		shelterNeedBeforeUpdateHooks = append(shelterNeedBeforeUpdateHooks, shelterNeedHook)
		shelterNeedBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		shelterNeedAfterUpdateMu.Lock()
		shelterNeedAfterUpdateHooks = append(shelterNeedAfterUpdateHooks, shelterNeedHook)
		shelterNeedAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		shelterNeedBeforeDeleteMu.Lock()
		shelterNeedBeforeDeleteHooks = append(shelterNeedBeforeDeleteHooks, shelterNeedHook)
		shelterNeedBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		shelterNeedAfterDeleteMu.Lock()
		shelterNeedAfterDeleteHooks = append(shelterNeedAfterDeleteHooks, shelterNeedHook)
		shelterNeedAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		shelterNeedBeforeUpsertMu.Lock()
		shelterNeedBeforeUpsertHooks = append(shelterNeedBeforeUpsertHooks, shelterNeedHook)
		shelterNeedBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		shelterNeedAfterUpsertMu.Lock()
		shelterNeedAfterUpsertHooks = append(shelterNeedAfterUpsertHooks, shelterNeedHook)
		shelterNeedAfterUpsertMu.Unlock()
	}
}

// One returns a single shelterNeed record from the query.
func (q shelterNeedQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ShelterNeed, error) {
	o := &ShelterNeed{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for shelter_needs")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ShelterNeed records from the query.
func (q shelterNeedQuery) All(ctx context.Context, exec boil.ContextExecutor) (ShelterNeedSlice, error) {
	var o []*ShelterNeed

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ShelterNeed slice")
	}

	if len(shelterNeedAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ShelterNeed records in the query.
func (q shelterNeedQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count shelter_needs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q shelterNeedQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if shelter_needs exists")
	}

	return count > 0, nil
}

// Shelter pointed to by the foreign key.
func (o *ShelterNeed) Shelter(mods ...qm.QueryMod) shelterQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ShelterID),
	}

	queryMods = append(queryMods, mods...)

	return Shelters(queryMods...)
}

// LoadShelter allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (shelterNeedL) LoadShelter(ctx context.Context, e boil.ContextExecutor, singular bool, maybeShelterNeed interface{}, mods queries.Applicator) error {
	var slice []*ShelterNeed
	var object *ShelterNeed

	if singular {
		var ok bool
		object, ok = maybeShelterNeed.(*ShelterNeed)
		if !ok {
			object = new(ShelterNeed)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeShelterNeed)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeShelterNeed))
			}
		}
	} else {
		s, ok := maybeShelterNeed.(*[]*ShelterNeed)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeShelterNeed)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeShelterNeed))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &shelterNeedR{}
		}
		args[object.ShelterID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &shelterNeedR{}
			}

			args[obj.ShelterID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`shelter`),
		qm.WhereIn(`shelter.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Shelter")
	}

	var resultSlice []*Shelter
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Shelter")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for shelter")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for shelter")
	}

	if len(shelterAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Shelter = foreign
		if foreign.R == nil {
			foreign.R = &shelterR{}
		}
		foreign.R.ShelterNeeds = append(foreign.R.ShelterNeeds, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ShelterID == foreign.ID {
				local.R.Shelter = foreign
				if foreign.R == nil {
					foreign.R = &shelterR{}
				}
				foreign.R.ShelterNeeds = append(foreign.R.ShelterNeeds, local)
				break
			}
		}
	}

	return nil
}

// SetShelter of the shelterNeed to the related item.
// Sets o.R.Shelter to related.
// Adds o to related.R.ShelterNeeds.
func (o *ShelterNeed) SetShelter(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Shelter) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"shelter_needs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"shelter_id"}),
		strmangle.WhereClause("\"", "\"", 2, shelterNeedPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ShelterID = related.ID
	if o.R == nil {
		o.R = &shelterNeedR{
			Shelter: related,
		}
	} else {
		o.R.Shelter = related
	}

	if related.R == nil {
		related.R = &shelterR{
			ShelterNeeds: ShelterNeedSlice{o},
		}
	} else {
		related.R.ShelterNeeds = append(related.R.ShelterNeeds, o)
	}

	return nil
}

// ShelterNeeds retrieves all the records using an executor.
func ShelterNeeds(mods ...qm.QueryMod) shelterNeedQuery {
	mods = append(mods, qm.From("\"shelter_needs\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"shelter_needs\".*"})
	}

	return shelterNeedQuery{q}
}

// FindShelterNeed retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindShelterNeed(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ShelterNeed, error) {
	shelterNeedObj := &ShelterNeed{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"shelter_needs\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, shelterNeedObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from shelter_needs")
	}

	if err = shelterNeedObj.doAfterSelectHooks(ctx, exec); err != nil {
		return shelterNeedObj, err
	}

	return shelterNeedObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ShelterNeed) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no shelter_needs provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(shelterNeedColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	shelterNeedInsertCacheMut.RLock()
	cache, cached := shelterNeedInsertCache[key]
	shelterNeedInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			shelterNeedAllColumns,
			shelterNeedColumnsWithDefault,
			shelterNeedColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(shelterNeedType, shelterNeedMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(shelterNeedType, shelterNeedMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"shelter_needs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"shelter_needs\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into shelter_needs")
	}

	if !cached {
		shelterNeedInsertCacheMut.Lock()
		shelterNeedInsertCache[key] = cache
		shelterNeedInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ShelterNeed.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ShelterNeed) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	shelterNeedUpdateCacheMut.RLock()
	cache, cached := shelterNeedUpdateCache[key]
	shelterNeedUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			shelterNeedAllColumns,
			shelterNeedPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update shelter_needs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"shelter_needs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, shelterNeedPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(shelterNeedType, shelterNeedMapping, append(wl, shelterNeedPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update shelter_needs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for shelter_needs")
	}

	if !cached {
		shelterNeedUpdateCacheMut.Lock()
		shelterNeedUpdateCache[key] = cache
		shelterNeedUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q shelterNeedQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for shelter_needs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for shelter_needs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ShelterNeedSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shelterNeedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"shelter_needs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, shelterNeedPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in shelterNeed slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all shelterNeed")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ShelterNeed) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no shelter_needs provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(shelterNeedColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	shelterNeedUpsertCacheMut.RLock()
	cache, cached := shelterNeedUpsertCache[key]
	shelterNeedUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			shelterNeedAllColumns,
			shelterNeedColumnsWithDefault,
			shelterNeedColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			shelterNeedAllColumns,
			shelterNeedPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert shelter_needs, could not build update column list")
		}

		ret := strmangle.SetComplement(shelterNeedAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(shelterNeedPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert shelter_needs, could not build conflict column list")
			}

			conflict = make([]string, len(shelterNeedPrimaryKeyColumns))
			copy(conflict, shelterNeedPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"shelter_needs\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(shelterNeedType, shelterNeedMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(shelterNeedType, shelterNeedMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert shelter_needs")
	}

	if !cached {
		shelterNeedUpsertCacheMut.Lock()
		shelterNeedUpsertCache[key] = cache
		shelterNeedUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ShelterNeed record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ShelterNeed) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ShelterNeed provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), shelterNeedPrimaryKeyMapping)
	sql := "DELETE FROM \"shelter_needs\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from shelter_needs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for shelter_needs")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q shelterNeedQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no shelterNeedQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from shelter_needs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for shelter_needs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ShelterNeedSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(shelterNeedBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shelterNeedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"shelter_needs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, shelterNeedPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from shelterNeed slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for shelter_needs")
	}

	if len(shelterNeedAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ShelterNeed) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindShelterNeed(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ShelterNeedSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ShelterNeedSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shelterNeedPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"shelter_needs\".* FROM \"shelter_needs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, shelterNeedPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ShelterNeedSlice")
	}

	*o = slice

	return nil
}

// ShelterNeedExists checks if the ShelterNeed row exists.
func ShelterNeedExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"shelter_needs\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if shelter_needs exists")
	}

	return exists, nil
}

// Exists checks if the ShelterNeed row exists.
func (o *ShelterNeed) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ShelterNeedExists(ctx, exec, o.ID)
}
