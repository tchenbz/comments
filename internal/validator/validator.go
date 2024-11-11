package validator

import (
    "slices"
  )

 
// We will create a new type named Validator
type Validator struct {
    Errors map[string]string
} 

// Construct a new Validator and return a pointer to it
// All validation errors go into this one Validator instance
func New() *Validator {
    return &Validator {
        Errors: make(map[string]string),
    }
}

// Let's check  to see if the Validator's map contains any entries 
func (v *Validator) IsEmpty() bool {
    return len(v.Errors) == 0
}

// Add a new error entry to the Validator's error map
// Check first if an entry with the same key does not already exist
func (v *Validator) AddError(key string, message string) {
    _, exists := v.Errors[key]
    if !exists {
        v.Errors[key] = message
    }
}

// If any validation check returns false, then we will
// make an entry into our Validator's error map 
func (v *Validator) Check(acceptable bool, key string, message string) {
    if !acceptable {
       v.AddError(key, message)
    }
}

// Check for permitted values
func PermittedValue(value string, permittedValues ...string) bool {
    return slices.Contains(permittedValues, value)  
}

