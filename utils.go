/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package pdf

import (
	"fmt"
	"sort"

	
)

// Inspect analyzes the document object structure.
func (this *PdfReader) Inspect() (map[string]int, error) {
	return this.parser.inspect()
}

func getUniDocVersion() string {
	return Version
}

/*
 * Inspect object types.
 * Go through all objects in the cross ref table and detect the types.
 */
func (this *PdfParser) inspect() (map[string]int, error) {
	Log.Debug("--------INSPECT ----------")
	Log.Debug("Xref table:")

	objTypes := map[string]int{}
	objCount := 0
	failedCount := 0

	keys := []int{}
	for k, _ := range this.xrefs {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	i := 0
	for _, k := range keys {
		xref := this.xrefs[k]
		if xref.objectNumber == 0 {
			continue
		}
		objCount++
		Log.Debug("==========")
		Log.Debug("Looking up object number: %d", xref.objectNumber)
		o, err := this.LookupByNumber(xref.objectNumber)
		if err != nil {
			Log.Debug("ERROR: Fail to lookup obj %d (%s)", xref.objectNumber, err)
			failedCount++
			continue
		}

		Log.Debug("obj: %s", o)

		iobj, isIndirect := o.(*PdfIndirectObject)
		if isIndirect {
			Log.Debug("IND OOBJ %d: %s", xref.objectNumber, iobj)
			dict, isDict := iobj.PdfObject.(*PdfObjectDictionary)
			if isDict {
				// Check if has Type parameter.
				if ot, has := (*dict)["Type"].(*PdfObjectName); has {
					otype := string(*ot)
					Log.Debug("---> Obj type: %s", otype)
					_, isDefined := objTypes[otype]
					if isDefined {
						objTypes[otype]++
					} else {
						objTypes[otype] = 1
					}
				} else if ot, has := (*dict)["Subtype"].(*PdfObjectName); has {
					// Check if subtype
					otype := string(*ot)
					Log.Debug("---> Obj subtype: %s", otype)
					_, isDefined := objTypes[otype]
					if isDefined {
						objTypes[otype]++
					} else {
						objTypes[otype] = 1
					}
				}
				if val, has := (*dict)["S"].(*PdfObjectName); has && *val == "JavaScript" {
					// Check if Javascript.
					_, isDefined := objTypes["JavaScript"]
					if isDefined {
						objTypes["JavaScript"]++
					} else {
						objTypes["JavaScript"] = 1
					}
				}

			}
		} else if sobj, isStream := o.(*PdfObjectStream); isStream {
			if otype, ok := (*(sobj.PdfObjectDictionary))["Type"].(*PdfObjectName); ok {
				Log.Debug("--> Stream object type: %s", *otype)
				k := string(*otype)
				if _, isDefined := objTypes[k]; isDefined {
					objTypes[k]++
				} else {
					objTypes[k] = 1
				}
			}
		} else { // Direct.
			dict, isDict := o.(*PdfObjectDictionary)
			if isDict {
				ot, isName := (*dict)["Type"].(*PdfObjectName)
				if isName {
					otype := string(*ot)
					Log.Debug("--- obj type %s", otype)
					objTypes[otype]++
				}
			}
			Log.Debug("DIRECT OBJ %d: %s", xref.objectNumber, o)
		}

		i++
	}
	Log.Debug("--------EOF INSPECT ----------")
	Log.Debug("=======")
	Log.Debug("Object count: %d", objCount)
	Log.Debug("Failed lookup: %d", failedCount)
	for t, c := range objTypes {
		Log.Debug("%s: %d", t, c)
	}
	Log.Debug("=======")

	if len(this.xrefs) < 1 {
		Log.Debug("ERROR: This document is invalid (xref table missing!)")
		return nil, fmt.Errorf("Invalid document (xref table missing)")
	}

	fontObjs, ok := objTypes["Font"]
	if !ok || fontObjs < 2 {
		Log.Debug("This document is probably scanned!")
	} else {
		Log.Debug("This document is valid for extraction!")
	}

	return objTypes, nil
}
