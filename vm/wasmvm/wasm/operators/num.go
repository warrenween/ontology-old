/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package operators

import (
	"github.com/ontio/ontology/vm/wasmvm/wasm"
)

var (
	I32Clz      = newOp(0x67, "i32.clz", []wasm.ValueType{wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I32Ctz      = newOp(0x68, "i32.ctz", []wasm.ValueType{wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I32Popcnt   = newOp(0x69, "i32.popcnt", []wasm.ValueType{wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I32Add      = newOp(0x6a, "i32.add", []wasm.ValueType{wasm.ValueTypeI32, wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I32Sub      = newOp(0x6b, "i32.sub", []wasm.ValueType{wasm.ValueTypeI32, wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I32Mul      = newOp(0x6c, "i32.mul", []wasm.ValueType{wasm.ValueTypeI32, wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I32DivS     = newOp(0x6d, "i32.div_s", []wasm.ValueType{wasm.ValueTypeI32, wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I32DivU     = newOp(0x6e, "i32.div_u", []wasm.ValueType{wasm.ValueTypeI32, wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I32RemS     = newOp(0x6f, "i32.rem_s", []wasm.ValueType{wasm.ValueTypeI32, wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I32RemU     = newOp(0x70, "i32.rem_u", []wasm.ValueType{wasm.ValueTypeI32, wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I32And      = newOp(0x71, "i32.and", []wasm.ValueType{wasm.ValueTypeI32, wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I32Or       = newOp(0x72, "i32.or", []wasm.ValueType{wasm.ValueTypeI32, wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I32Xor      = newOp(0x73, "i32.xor", []wasm.ValueType{wasm.ValueTypeI32, wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I32Shl      = newOp(0x74, "i32.shl", []wasm.ValueType{wasm.ValueTypeI32, wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I32ShrS     = newOp(0x75, "i32.shr_s", []wasm.ValueType{wasm.ValueTypeI32, wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I32ShrU     = newOp(0x76, "i32.shr_u", []wasm.ValueType{wasm.ValueTypeI32, wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I32Rotl     = newOp(0x77, "i32.rotl", []wasm.ValueType{wasm.ValueTypeI32, wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I32Rotr     = newOp(0x78, "i32.rotr", []wasm.ValueType{wasm.ValueTypeI32, wasm.ValueTypeI32}, wasm.ValueTypeI32)
	I64Clz      = newOp(0x79, "i64.clz", []wasm.ValueType{wasm.ValueTypeI64}, wasm.ValueTypeI64)
	I64Ctz      = newOp(0x7a, "i64.ctz", []wasm.ValueType{wasm.ValueTypeI64}, wasm.ValueTypeI64)
	I64Popcnt   = newOp(0x7b, "i64.popcnt", []wasm.ValueType{wasm.ValueTypeI64}, wasm.ValueTypeI64)
	I64Add      = newOp(0x7c, "i64.add", []wasm.ValueType{wasm.ValueTypeI64, wasm.ValueTypeI64}, wasm.ValueTypeI64)
	I64Sub      = newOp(0x7d, "i64.sub", []wasm.ValueType{wasm.ValueTypeI64, wasm.ValueTypeI64}, wasm.ValueTypeI64)
	I64Mul      = newOp(0x7e, "i64.mul", []wasm.ValueType{wasm.ValueTypeI64, wasm.ValueTypeI64}, wasm.ValueTypeI64)
	I64DivS     = newOp(0x7f, "i64.div_s", []wasm.ValueType{wasm.ValueTypeI64, wasm.ValueTypeI64}, wasm.ValueTypeI64)
	I64DivU     = newOp(0x80, "i64.div_u", []wasm.ValueType{wasm.ValueTypeI64, wasm.ValueTypeI64}, wasm.ValueTypeI64)
	I64RemS     = newOp(0x81, "i64.div_u", []wasm.ValueType{wasm.ValueTypeI64, wasm.ValueTypeI64}, wasm.ValueTypeI64)
	I64RemU     = newOp(0x82, "i64.rem_u", []wasm.ValueType{wasm.ValueTypeI64, wasm.ValueTypeI64}, wasm.ValueTypeI64)
	I64And      = newOp(0x83, "i64.and", []wasm.ValueType{wasm.ValueTypeI64, wasm.ValueTypeI64}, wasm.ValueTypeI64)
	I64Or       = newOp(0x84, "i64.or", []wasm.ValueType{wasm.ValueTypeI64, wasm.ValueTypeI64}, wasm.ValueTypeI64)
	I64Xor      = newOp(0x85, "i64.xor", []wasm.ValueType{wasm.ValueTypeI64, wasm.ValueTypeI64}, wasm.ValueTypeI64)
	I64Shl      = newOp(0x86, "i64.shl", []wasm.ValueType{wasm.ValueTypeI64, wasm.ValueTypeI64}, wasm.ValueTypeI64)
	I64ShrS     = newOp(0x87, "i64.shr_s", []wasm.ValueType{wasm.ValueTypeI64, wasm.ValueTypeI64}, wasm.ValueTypeI64)
	I64ShrU     = newOp(0x88, "i64.shr_u", []wasm.ValueType{wasm.ValueTypeI64, wasm.ValueTypeI64}, wasm.ValueTypeI64)
	I64Rotl     = newOp(0x89, "i64.rotl", []wasm.ValueType{wasm.ValueTypeI64, wasm.ValueTypeI64}, wasm.ValueTypeI64)
	I64Rotr     = newOp(0x8a, "i64.rotr", []wasm.ValueType{wasm.ValueTypeI64, wasm.ValueTypeI64}, wasm.ValueTypeI64)
	F32Abs      = newOp(0x8b, "f32.abs", []wasm.ValueType{wasm.ValueTypeF32}, wasm.ValueTypeF32)
	F32Neg      = newOp(0x8c, "f32.neg", []wasm.ValueType{wasm.ValueTypeF32}, wasm.ValueTypeF32)
	F32Ceil     = newOp(0x8d, "f32.ceil", []wasm.ValueType{wasm.ValueTypeF32}, wasm.ValueTypeF32)
	F32Floor    = newOp(0x8e, "f32.floor", []wasm.ValueType{wasm.ValueTypeF32}, wasm.ValueTypeF32)
	F32Trunc    = newOp(0x8f, "f32.trunc", []wasm.ValueType{wasm.ValueTypeF32}, wasm.ValueTypeF32)
	F32Nearest  = newOp(0x90, "f32.nearest", []wasm.ValueType{wasm.ValueTypeF32}, wasm.ValueTypeF32)
	F32Sqrt     = newOp(0x91, "f32.sqrt", []wasm.ValueType{wasm.ValueTypeF32}, wasm.ValueTypeF32)
	F32Add      = newOp(0x92, "f32.add", []wasm.ValueType{wasm.ValueTypeF32, wasm.ValueTypeF32}, wasm.ValueTypeF32)
	F32Sub      = newOp(0x93, "f32.sub", []wasm.ValueType{wasm.ValueTypeF32, wasm.ValueTypeF32}, wasm.ValueTypeF32)
	F32Mul      = newOp(0x94, "f32.mul", []wasm.ValueType{wasm.ValueTypeF32, wasm.ValueTypeF32}, wasm.ValueTypeF32)
	F32Div      = newOp(0x95, "f32.div", []wasm.ValueType{wasm.ValueTypeF32, wasm.ValueTypeF32}, wasm.ValueTypeF32)
	F32Min      = newOp(0x96, "f32.min", []wasm.ValueType{wasm.ValueTypeF32, wasm.ValueTypeF32}, wasm.ValueTypeF32)
	F32Max      = newOp(0x97, "f32.max", []wasm.ValueType{wasm.ValueTypeF32, wasm.ValueTypeF32}, wasm.ValueTypeF32)
	F32Copysign = newOp(0x98, "f32.copysign", []wasm.ValueType{wasm.ValueTypeF32, wasm.ValueTypeF32}, wasm.ValueTypeF32)
	F64Abs      = newOp(0x99, "f64.abs", []wasm.ValueType{wasm.ValueTypeF64}, wasm.ValueTypeF64)
	F64Neg      = newOp(0x9a, "f64.neg", []wasm.ValueType{wasm.ValueTypeF64}, wasm.ValueTypeF64)
	F64Ceil     = newOp(0x9b, "f64.ceil", []wasm.ValueType{wasm.ValueTypeF64}, wasm.ValueTypeF64)
	F64Floor    = newOp(0x9c, "f64.floor", []wasm.ValueType{wasm.ValueTypeF64}, wasm.ValueTypeF64)
	F64Trunc    = newOp(0x9d, "f64.trunc", []wasm.ValueType{wasm.ValueTypeF64}, wasm.ValueTypeF64)
	F64Nearest  = newOp(0x9e, "f64.nearest", []wasm.ValueType{wasm.ValueTypeF64}, wasm.ValueTypeF64)
	F64Sqrt     = newOp(0x9f, "f64.sqrt", []wasm.ValueType{wasm.ValueTypeF64}, wasm.ValueTypeF64)
	F64Add      = newOp(0xa0, "f64.add", []wasm.ValueType{wasm.ValueTypeF64, wasm.ValueTypeF64}, wasm.ValueTypeF64)
	F64Sub      = newOp(0xa1, "f64.sub", []wasm.ValueType{wasm.ValueTypeF64, wasm.ValueTypeF64}, wasm.ValueTypeF64)
	F64Mul      = newOp(0xa2, "f64.mul", []wasm.ValueType{wasm.ValueTypeF64, wasm.ValueTypeF64}, wasm.ValueTypeF64)
	F64Div      = newOp(0xa3, "f64.div", []wasm.ValueType{wasm.ValueTypeF64, wasm.ValueTypeF64}, wasm.ValueTypeF64)
	F64Min      = newOp(0xa4, "f64.min", []wasm.ValueType{wasm.ValueTypeF64, wasm.ValueTypeF64}, wasm.ValueTypeF64)
	F64Max      = newOp(0xa5, "f64.max", []wasm.ValueType{wasm.ValueTypeF64, wasm.ValueTypeF64}, wasm.ValueTypeF64)
	F64Copysign = newOp(0xa6, "f64.copysign", []wasm.ValueType{wasm.ValueTypeF64, wasm.ValueTypeF64}, wasm.ValueTypeF64)
)
