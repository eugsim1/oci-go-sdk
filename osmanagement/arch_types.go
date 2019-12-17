// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// OSMS
//
// OS Management as a Service API definition
//

package osmanagement

// ArchTypesEnum Enum with underlying type: string
type ArchTypesEnum string

// Set of constants representing the allowable values for ArchTypesEnum
const (
	ArchTypesIa32        ArchTypesEnum = "IA_32"
	ArchTypesX8664       ArchTypesEnum = "X86_64"
	ArchTypesAarch64     ArchTypesEnum = "AARCH64"
	ArchTypesSparc       ArchTypesEnum = "SPARC"
	ArchTypesAmd64Debian ArchTypesEnum = "AMD64_DEBIAN"
)

var mappingArchTypes = map[string]ArchTypesEnum{
	"IA_32":        ArchTypesIa32,
	"X86_64":       ArchTypesX8664,
	"AARCH64":      ArchTypesAarch64,
	"SPARC":        ArchTypesSparc,
	"AMD64_DEBIAN": ArchTypesAmd64Debian,
}

// GetArchTypesEnumValues Enumerates the set of values for ArchTypesEnum
func GetArchTypesEnumValues() []ArchTypesEnum {
	values := make([]ArchTypesEnum, 0)
	for _, v := range mappingArchTypes {
		values = append(values, v)
	}
	return values
}
