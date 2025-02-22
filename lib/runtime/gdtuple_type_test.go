/*
 * Copyright (C) 2023 The GDLang Team.
 *
 * This file is part of GDLang.
 *
 * GDLang is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * GDLang is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with GDLang.  If not, see <http://www.gnu.org/licenses/>.
 */

package runtime_test

import (
	"gdlang/lib/runtime"
	"testing"
)

func TestSimpleTupleType(t *testing.T) {
	tupleType := runtime.NewGDTupleType(runtime.GDIntType, runtime.GDStringType)

	if len(tupleType) != 2 {
		t.Errorf("Expected 2 elements, got %v", len(tupleType))
	}

	if tupleType[0].GetCode() != runtime.GDIntTypeCode {
		t.Errorf("Expected number type, got %v", tupleType[0].GetCode())
	}

	if tupleType[1].GetCode() != runtime.GDStringTypeCode {
		t.Errorf("Expected string type, got %v", tupleType[1].GetCode())
	}

	tupleTypeStr := tupleType.ToString()
	if tupleTypeStr != "(int, string)" {
		t.Errorf("Expected number, string, got %v", tupleTypeStr)
	}
}

func TestNestedTupleType(t *testing.T) {
	tupleType := runtime.NewGDTupleType(runtime.GDIntType, runtime.NewGDTupleType(runtime.GDIntType, runtime.GDStringType))

	if len(tupleType) != 2 {
		t.Errorf("Expected 2 elements, got %v", len(tupleType))
	}

	if tupleType[0].GetCode() != runtime.GDIntTypeCode {
		t.Errorf("Expected number type, got %v", tupleType[0].GetCode())
	}

	if tupleType[1].GetCode() != runtime.GDTupleTypeCode {
		t.Errorf("Expected tuple type, got %v", tupleType[1].GetCode())
	}

	tupleTypeStr := tupleType.ToString()
	if tupleTypeStr != "(int, (int, string))" {
		t.Errorf("Expected (int, (int, string)), got %v", tupleTypeStr)
	}
}
