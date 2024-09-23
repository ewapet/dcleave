# DCleave

## Overview

A helper library for working with abstracted functionality that isn't known at compile-time. 

It enables defining "translators", which retrieve data or execute behavior without affecting the underlying model.

**Note**: This shouldn't be necessary for everyday use cases. Direct dependencies on this library are not recommended. 
This component is a spin-off from a larger project.

```go
orgTranslator := dcleave.NewTranslator(...)

orgModel := organization.GetID(42)
newUser := user.New(...)


orgTranslator.Put(orgModel, newUser, []any{"users", "quarantined"})
```