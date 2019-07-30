package spin

// MongoDBSpinner is a spinner to spin up a dev mongodb container
type MongoDBSpinner struct{}

func (s *MongoDBSpinner) Spin(c *RunConfig) error {
	return nil
}
