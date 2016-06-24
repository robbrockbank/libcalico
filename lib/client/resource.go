package client

// Save the resource in the datastore:
// -  can_create is used to indicate whether the resource can be created if it does not exist
// -  can_replace is used to indicate whether the resource can be updated if it already exists
//
// The Resource may be of Kind "list" and therefore contain zero or more resources to save.
// If a list of resources is specified, they are saved in the list order, and this function
// returns after saving all resources, or after hitting an error.  The function returns the
// number of resources successfully updated.
func (c *CalicoClient) SaveResource(r interface{}, canCreate, canReplace bool) (int, error) {
	return 0, nil
}

// Load the resource(s) from the datastore:
// -  The Resource does not need to contain the Spec - if specified it is ignored.
// -  If the supplied Resource metadata contains missing identifiers, the query will wildcard
//    those identifiers.  If multiple values are returned, the returned resource will be of
//    kind "list".
//
// The Resource may be of Kind "list" and therefore contain zero or more resources to save.
// If a list of resources is specified, they are saved in the list order, and this function
// returns after saving all resources, or after hitting an error.  The function returns the
// number of resources successfully updated.
func (c *CalicoClient) LoadResource(r interface{}) (interface{}, error) {
	return nil, nil
}

// Delete the resource(s) from the datastore:
// -  The Resource does not need to contain the Spec - if specified it is ignored.
// -  The Resource metadata should contain all identifiers required to uniquely identify a
//    single resource.
// -  The ignore_not_present flag indicates whether attempts to delete a missing resource is
//    ignored, or treated as an error.
//
// The Resource may be of Kind "list" and therefore contain zero or more resources to delete.
// If a list of resources is specified, they are deleted in the list order, and this function
// returns after saving all resources, or after hitting an error.  The function returns the
// number of resources successfully deleted.
func (c *CalicoClient) DeleteResource(r interface{}, ignoreNotPresent bool) (int, error) {
	return 0, nil
}