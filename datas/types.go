// This file was generated by nomdl/codegen.

package datas

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __datasPackageInFile_types_CachedRef = __datasPackageInFile_types_Ref()

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func __datasPackageInFile_types_Ref() ref.Ref {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("Commit",
			[]types.Field{
				types.Field{"value", types.MakePrimitiveTypeRef(types.ValueKind), false},
				types.Field{"parents", types.MakeCompoundTypeRef("", types.SetKind, types.MakeTypeRef(ref.Ref{}, 0)), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{})
	return types.RegisterPackage(&p)
}

// Commit

type Commit struct {
	m   types.Map
	ref *ref.Ref
}

func NewCommit() Commit {
	return Commit{types.NewMap(
		types.NewString("$type"), types.MakeTypeRef(__datasPackageInFile_types_CachedRef, 0),
		types.NewString("value"), types.Bool(false),
		types.NewString("parents"), NewSetOfCommit(),
	), &ref.Ref{}}
}

var __typeRefForCommit = types.MakeTypeRef(__datasPackageInFile_types_CachedRef, 0)

func (m Commit) TypeRef() types.TypeRef {
	return __typeRefForCommit
}

func init() {
	types.RegisterFromValFunction(__typeRefForCommit, func(v types.Value) types.Value {
		return CommitFromVal(v)
	})
}

func CommitFromVal(val types.Value) Commit {
	// TODO: Do we still need FromVal?
	if val, ok := val.(Commit); ok {
		return val
	}
	// TODO: Validate here
	return Commit{val.(types.Map), &ref.Ref{}}
}

func (s Commit) NomsValue() types.Value {
	// TODO: Remove this
	return s
}

func (s Commit) InternalImplementation() types.Map {
	return s.m
}

func (s Commit) Equals(other types.Value) bool {
	if other, ok := other.(Commit); ok {
		return s.Ref() == other.Ref()
	}
	return false
}

func (s Commit) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Commit) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s Commit) Value() types.Value {
	return s.m.Get(types.NewString("value"))
}

func (s Commit) SetValue(val types.Value) Commit {
	return Commit{s.m.Set(types.NewString("value"), val), &ref.Ref{}}
}

func (s Commit) Parents() SetOfCommit {
	return s.m.Get(types.NewString("parents")).(SetOfCommit)
}

func (s Commit) SetParents(val SetOfCommit) Commit {
	return Commit{s.m.Set(types.NewString("parents"), val), &ref.Ref{}}
}

// MapOfStringToCommit

type MapOfStringToCommit struct {
	m   types.Map
	ref *ref.Ref
}

func NewMapOfStringToCommit() MapOfStringToCommit {
	return MapOfStringToCommit{types.NewMap(), &ref.Ref{}}
}

func MapOfStringToCommitFromVal(val types.Value) MapOfStringToCommit {
	// TODO: Do we still need FromVal?
	if val, ok := val.(MapOfStringToCommit); ok {
		return val
	}
	// TODO: Validate here
	return MapOfStringToCommit{val.(types.Map), &ref.Ref{}}
}

func (m MapOfStringToCommit) NomsValue() types.Value {
	// TODO: Remove this
	return m
}

func (m MapOfStringToCommit) InternalImplementation() types.Map {
	return m.m
}

func (m MapOfStringToCommit) Equals(other types.Value) bool {
	if other, ok := other.(MapOfStringToCommit); ok {
		return m.Ref() == other.Ref()
	}
	return false
}

func (m MapOfStringToCommit) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfStringToCommit) Chunks() (futures []types.Future) {
	futures = append(futures, m.TypeRef().Chunks()...)
	futures = append(futures, m.m.Chunks()...)
	return
}

// A Noms Value that describes MapOfStringToCommit.
var __typeRefForMapOfStringToCommit types.TypeRef

func (m MapOfStringToCommit) TypeRef() types.TypeRef {
	return __typeRefForMapOfStringToCommit
}

func init() {
	__typeRefForMapOfStringToCommit = types.MakeCompoundTypeRef("", types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeTypeRef(__datasPackageInFile_types_CachedRef, 0))
	types.RegisterFromValFunction(__typeRefForMapOfStringToCommit, func(v types.Value) types.Value {
		return MapOfStringToCommitFromVal(v)
	})
}

func (m MapOfStringToCommit) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToCommit) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToCommit) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToCommit) Get(p string) Commit {
	return m.m.Get(types.NewString(p)).(Commit)
}

func (m MapOfStringToCommit) MaybeGet(p string) (Commit, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return NewCommit(), false
	}
	return v.(Commit), ok
}

func (m MapOfStringToCommit) Set(k string, v Commit) MapOfStringToCommit {
	return MapOfStringToCommit{m.m.Set(types.NewString(k), v), &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfStringToCommit) Remove(p string) MapOfStringToCommit {
	return MapOfStringToCommit{m.m.Remove(types.NewString(p)), &ref.Ref{}}
}

type MapOfStringToCommitIterCallback func(k string, v Commit) (stop bool)

func (m MapOfStringToCommit) Iter(cb MapOfStringToCommitIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(Commit))
	})
}

type MapOfStringToCommitIterAllCallback func(k string, v Commit)

func (m MapOfStringToCommit) IterAll(cb MapOfStringToCommitIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v.(Commit))
	})
}

type MapOfStringToCommitFilterCallback func(k string, v Commit) (keep bool)

func (m MapOfStringToCommit) Filter(cb MapOfStringToCommitFilterCallback) MapOfStringToCommit {
	nm := NewMapOfStringToCommit()
	m.IterAll(func(k string, v Commit) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}

// SetOfCommit

type SetOfCommit struct {
	s   types.Set
	ref *ref.Ref
}

func NewSetOfCommit() SetOfCommit {
	return SetOfCommit{types.NewSet(), &ref.Ref{}}
}

func SetOfCommitFromVal(val types.Value) SetOfCommit {
	// TODO: Do we still need FromVal?
	if val, ok := val.(SetOfCommit); ok {
		return val
	}
	return SetOfCommit{val.(types.Set), &ref.Ref{}}
}

func (s SetOfCommit) NomsValue() types.Value {
	// TODO: Remove this
	return s
}

func (s SetOfCommit) InternalImplementation() types.Set {
	return s.s
}

func (s SetOfCommit) Equals(other types.Value) bool {
	if other, ok := other.(SetOfCommit); ok {
		return s.Ref() == other.Ref()
	}
	return false
}

func (s SetOfCommit) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s SetOfCommit) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.s.Chunks()...)
	return
}

// A Noms Value that describes SetOfCommit.
var __typeRefForSetOfCommit types.TypeRef

func (m SetOfCommit) TypeRef() types.TypeRef {
	return __typeRefForSetOfCommit
}

func init() {
	__typeRefForSetOfCommit = types.MakeCompoundTypeRef("", types.SetKind, types.MakeTypeRef(__datasPackageInFile_types_CachedRef, 0))
	types.RegisterFromValFunction(__typeRefForSetOfCommit, func(v types.Value) types.Value {
		return SetOfCommitFromVal(v)
	})
}

func (s SetOfCommit) Empty() bool {
	return s.s.Empty()
}

func (s SetOfCommit) Len() uint64 {
	return s.s.Len()
}

func (s SetOfCommit) Has(p Commit) bool {
	return s.s.Has(p)
}

type SetOfCommitIterCallback func(p Commit) (stop bool)

func (s SetOfCommit) Iter(cb SetOfCommitIterCallback) {
	s.s.Iter(func(v types.Value) bool {
		return cb(v.(Commit))
	})
}

type SetOfCommitIterAllCallback func(p Commit)

func (s SetOfCommit) IterAll(cb SetOfCommitIterAllCallback) {
	s.s.IterAll(func(v types.Value) {
		cb(v.(Commit))
	})
}

type SetOfCommitFilterCallback func(p Commit) (keep bool)

func (s SetOfCommit) Filter(cb SetOfCommitFilterCallback) SetOfCommit {
	ns := NewSetOfCommit()
	s.IterAll(func(v Commit) {
		if cb(v) {
			ns = ns.Insert(v)
		}
	})
	return ns
}

func (s SetOfCommit) Insert(p ...Commit) SetOfCommit {
	return SetOfCommit{s.s.Insert(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfCommit) Remove(p ...Commit) SetOfCommit {
	return SetOfCommit{s.s.Remove(s.fromElemSlice(p)...), &ref.Ref{}}
}

func (s SetOfCommit) Union(others ...SetOfCommit) SetOfCommit {
	return SetOfCommit{s.s.Union(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfCommit) Subtract(others ...SetOfCommit) SetOfCommit {
	return SetOfCommit{s.s.Subtract(s.fromStructSlice(others)...), &ref.Ref{}}
}

func (s SetOfCommit) Any() Commit {
	return s.s.Any().(Commit)
}

func (s SetOfCommit) fromStructSlice(p []SetOfCommit) []types.Set {
	r := make([]types.Set, len(p))
	for i, v := range p {
		r[i] = v.s
	}
	return r
}

func (s SetOfCommit) fromElemSlice(p []Commit) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}
