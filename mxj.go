// mxj - A collection of map[string]interface{} and associated XML and JSON utilities.
// Copyright 2012-2014 Charles Banning. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file

package mxj

import (
	"fmt"
	"strconv"
)

const (
	Cast         = true // for clarity - e.g., mxj.NewMapXml(doc, mxj.Cast)
	SafeEncoding = true // ditto - e.g., mv.Json(mxj.SafeEncoding)
)

type Map map[string]interface{}

// Allocate a Map.
func New() Map {
	m := make(map[string]interface{}, 0)
	return m
}

// Cast a Map to map[string]interface{}
func (mv Map) Old() map[string]interface{} {
	return mv
}

// --------------- StringIndent ... from x2j.WriteMap -------------

// Pretty print a Map.
func (mv Map) StringIndent(offset ...int) string {
	return writeMap(map[string]interface{}(mv), offset...)
}

// writeMap - dumps the map[string]interface{} for examination.
//	'offset' is initial indentation count; typically: Write(m).
func writeMap(m interface{}, offset ...int) string {
	var indent int
	if len(offset) == 1 {
		indent = offset[0]
	}

	var s string
	switch m.(type) {
	case nil:
		return "[nil] nil"
	case string:
		return "[string] " + m.(string)
	case float64:
		return "[float64] " + strconv.FormatFloat(m.(float64), 'e', 2, 64)
	case bool:
		return "[bool] " + strconv.FormatBool(m.(bool))
	case []interface{}:
		s += "[[]interface{}]"
		for i, v := range m.([]interface{}) {
			s += "\n"
			for i := 0; i < indent; i++ {
				s += "  "
			}
			s += "[item: " + strconv.FormatInt(int64(i), 10) + "]"
			switch v.(type) {
			case string, float64, bool:
				s += "\n"
			default:
				// noop
			}
			for i := 0; i < indent; i++ {
				s += "  "
			}
			s += writeMap(v, indent+1)
		}
	case map[string]interface{}:
		for k, v := range m.(map[string]interface{}) {
			s += "\n"
			for i := 0; i < indent; i++ {
				s += "  "
			}
			// s += "[map[string]interface{}] "+k+" :"+writeMap(v,indent+1)
			s += k + " :" + writeMap(v, indent+1)
		}
	default:
		// shouldn't ever be here ...
		s += fmt.Sprintf("unknown type for: %v", m)
	}
	return s
}
