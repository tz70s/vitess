/*
Copyright 2021 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by Sizegen. DO NOT EDIT.

package vindexes

import (
	"math"
	"reflect"
	"unsafe"
)

type cachedObject interface {
	CachedSize(alloc bool) int64
}

func (cached *AutoIncrement) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(48)
	}
	// field Column vitess.io/vitess/go/vt/sqlparser.ColIdent
	size += cached.Column.CachedSize(false)
	// field Sequence *vitess.io/vitess/go/vt/vtgate/vindexes.Table
	size += cached.Sequence.CachedSize(true)
	return size
}
func (cached *Binary) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(16)
	}
	// field name string
	size += int64(len(cached.name))
	return size
}
func (cached *BinaryMD5) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(16)
	}
	// field name string
	size += int64(len(cached.name))
	return size
}
func (cached *CFC) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(64)
	}
	// field name string
	size += int64(len(cached.name))
	// field offsets []int
	{
		size += int64(cap(cached.offsets)) * int64(8)
	}
	// field prefixVindex vitess.io/vitess/go/vt/vtgate/vindexes.SingleColumn
	if cc, ok := cached.prefixVindex.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *Column) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(44)
	}
	// field Name vitess.io/vitess/go/vt/sqlparser.ColIdent
	size += cached.Name.CachedSize(false)
	return size
}
func (cached *ColumnVindex) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(80)
	}
	// field Columns []vitess.io/vitess/go/vt/sqlparser.ColIdent
	{
		size += int64(cap(cached.Columns)) * int64(40)
		for _, elem := range cached.Columns {
			size += elem.CachedSize(false)
		}
	}
	// field Type string
	size += int64(len(cached.Type))
	// field Name string
	size += int64(len(cached.Name))
	// field Vindex vitess.io/vitess/go/vt/vtgate/vindexes.Vindex
	if cc, ok := cached.Vindex.(cachedObject); ok {
		size += cc.CachedSize(true)
	}
	return size
}
func (cached *ConsistentLookup) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(8)
	}
	// field clCommon *vitess.io/vitess/go/vt/vtgate/vindexes.clCommon
	size += cached.clCommon.CachedSize(true)
	return size
}
func (cached *ConsistentLookupUnique) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(8)
	}
	// field clCommon *vitess.io/vitess/go/vt/vtgate/vindexes.clCommon
	size += cached.clCommon.CachedSize(true)
	return size
}
func (cached *Hash) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(16)
	}
	// field name string
	size += int64(len(cached.name))
	return size
}
func (cached *Keyspace) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(17)
	}
	// field Name string
	size += int64(len(cached.Name))
	return size
}
func (cached *LookupHash) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(136)
	}
	// field name string
	size += int64(len(cached.name))
	// field lkp vitess.io/vitess/go/vt/vtgate/vindexes.lookupInternal
	size += cached.lkp.CachedSize(false)
	return size
}
func (cached *LookupHashUnique) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(136)
	}
	// field name string
	size += int64(len(cached.name))
	// field lkp vitess.io/vitess/go/vt/vtgate/vindexes.lookupInternal
	size += cached.lkp.CachedSize(false)
	return size
}
func (cached *LookupNonUnique) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(136)
	}
	// field name string
	size += int64(len(cached.name))
	// field lkp vitess.io/vitess/go/vt/vtgate/vindexes.lookupInternal
	size += cached.lkp.CachedSize(false)
	return size
}
func (cached *LookupUnicodeLooseMD5Hash) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(136)
	}
	// field name string
	size += int64(len(cached.name))
	// field lkp vitess.io/vitess/go/vt/vtgate/vindexes.lookupInternal
	size += cached.lkp.CachedSize(false)
	return size
}
func (cached *LookupUnicodeLooseMD5HashUnique) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(136)
	}
	// field name string
	size += int64(len(cached.name))
	// field lkp vitess.io/vitess/go/vt/vtgate/vindexes.lookupInternal
	size += cached.lkp.CachedSize(false)
	return size
}
func (cached *LookupUnique) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(136)
	}
	// field name string
	size += int64(len(cached.name))
	// field lkp vitess.io/vitess/go/vt/vtgate/vindexes.lookupInternal
	size += cached.lkp.CachedSize(false)
	return size
}
func (cached *Null) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(16)
	}
	// field name string
	size += int64(len(cached.name))
	return size
}
func (cached *Numeric) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(16)
	}
	// field name string
	size += int64(len(cached.name))
	return size
}

//go:nocheckptr
func (cached *NumericStaticMap) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(24)
	}
	// field name string
	size += int64(len(cached.name))
	// field lookup vitess.io/vitess/go/vt/vtgate/vindexes.NumericLookupTable
	if cached.lookup != nil {
		size += int64(48)
		hmap := reflect.ValueOf(cached.lookup)
		numBuckets := int(math.Pow(2, float64((*(*uint8)(unsafe.Pointer(hmap.Pointer() + uintptr(9)))))))
		numOldBuckets := (*(*uint16)(unsafe.Pointer(hmap.Pointer() + uintptr(10))))
		size += int64(numOldBuckets * 144)
		if len(cached.lookup) > 0 || numBuckets > 1 {
			size += int64(numBuckets * 144)
		}
	}
	return size
}
func (cached *RegionExperimental) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(24)
	}
	// field name string
	size += int64(len(cached.name))
	return size
}

//go:nocheckptr
func (cached *RegionJSON) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(32)
	}
	// field name string
	size += int64(len(cached.name))
	// field regionMap vitess.io/vitess/go/vt/vtgate/vindexes.RegionMap
	if cached.regionMap != nil {
		size += int64(48)
		hmap := reflect.ValueOf(cached.regionMap)
		numBuckets := int(math.Pow(2, float64((*(*uint8)(unsafe.Pointer(hmap.Pointer() + uintptr(9)))))))
		numOldBuckets := (*(*uint16)(unsafe.Pointer(hmap.Pointer() + uintptr(10))))
		size += int64(numOldBuckets * 208)
		if len(cached.regionMap) > 0 || numBuckets > 1 {
			size += int64(numBuckets * 208)
		}
		for k := range cached.regionMap {
			size += int64(len(k))
		}
	}
	return size
}
func (cached *ReverseBits) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(16)
	}
	// field name string
	size += int64(len(cached.name))
	return size
}
func (cached *Table) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(169)
	}
	// field Type string
	size += int64(len(cached.Type))
	// field Name vitess.io/vitess/go/vt/sqlparser.TableIdent
	size += cached.Name.CachedSize(false)
	// field Keyspace *vitess.io/vitess/go/vt/vtgate/vindexes.Keyspace
	size += cached.Keyspace.CachedSize(true)
	// field ColumnVindexes []*vitess.io/vitess/go/vt/vtgate/vindexes.ColumnVindex
	{
		size += int64(cap(cached.ColumnVindexes)) * int64(8)
		for _, elem := range cached.ColumnVindexes {
			size += elem.CachedSize(true)
		}
	}
	// field Ordered []*vitess.io/vitess/go/vt/vtgate/vindexes.ColumnVindex
	{
		size += int64(cap(cached.Ordered)) * int64(8)
		for _, elem := range cached.Ordered {
			size += elem.CachedSize(true)
		}
	}
	// field Owned []*vitess.io/vitess/go/vt/vtgate/vindexes.ColumnVindex
	{
		size += int64(cap(cached.Owned)) * int64(8)
		for _, elem := range cached.Owned {
			size += elem.CachedSize(true)
		}
	}
	// field AutoIncrement *vitess.io/vitess/go/vt/vtgate/vindexes.AutoIncrement
	size += cached.AutoIncrement.CachedSize(true)
	// field Columns []vitess.io/vitess/go/vt/vtgate/vindexes.Column
	{
		size += int64(cap(cached.Columns)) * int64(44)
		for _, elem := range cached.Columns {
			size += elem.CachedSize(false)
		}
	}
	// field Pinned []byte
	size += int64(cap(cached.Pinned))
	return size
}
func (cached *UnicodeLooseMD5) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(16)
	}
	// field name string
	size += int64(len(cached.name))
	return size
}
func (cached *UnicodeLooseXXHash) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(16)
	}
	// field name string
	size += int64(len(cached.name))
	return size
}
func (cached *XXHash) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(16)
	}
	// field name string
	size += int64(len(cached.name))
	return size
}
func (cached *clCommon) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(256)
	}
	// field name string
	size += int64(len(cached.name))
	// field lkp vitess.io/vitess/go/vt/vtgate/vindexes.lookupInternal
	size += cached.lkp.CachedSize(false)
	// field keyspace string
	size += int64(len(cached.keyspace))
	// field ownerTable string
	size += int64(len(cached.ownerTable))
	// field ownerColumns []string
	{
		size += int64(cap(cached.ownerColumns)) * int64(16)
		for _, elem := range cached.ownerColumns {
			size += int64(len(elem))
		}
	}
	// field lockLookupQuery string
	size += int64(len(cached.lockLookupQuery))
	// field lockOwnerQuery string
	size += int64(len(cached.lockOwnerQuery))
	// field insertLookupQuery string
	size += int64(len(cached.insertLookupQuery))
	// field updateLookupQuery string
	size += int64(len(cached.updateLookupQuery))
	return size
}
func (cached *lookupInternal) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(112)
	}
	// field Table string
	size += int64(len(cached.Table))
	// field FromColumns []string
	{
		size += int64(cap(cached.FromColumns)) * int64(16)
		for _, elem := range cached.FromColumns {
			size += int64(len(elem))
		}
	}
	// field To string
	size += int64(len(cached.To))
	// field sel string
	size += int64(len(cached.sel))
	// field ver string
	size += int64(len(cached.ver))
	// field del string
	size += int64(len(cached.del))
	return size
}
func (cached *prefixCFC) CachedSize(alloc bool) int64 {
	if cached == nil {
		return int64(0)
	}
	size := int64(0)
	if alloc {
		size += int64(24)
	}
	// field name string
	size += int64(len(cached.name))
	// field cfc *vitess.io/vitess/go/vt/vtgate/vindexes.CFC
	size += cached.cfc.CachedSize(true)
	return size
}
