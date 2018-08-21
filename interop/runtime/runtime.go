package runtime

// CheckWitness verifies if the given hash is the invoker of the contract.
func CheckWitness(hash []byte) bool {
	return true
}

// GetTrigger returns the smart contract invoke trigger which can be either
// verification or application.
func GetTrigger() byte {
	return 0x00
}

// Application returns the application trigger.
func Application() byte {
	return 0x10
}

// Verification returns the verification trigger.
func Verification() byte {
	return 0x00
}

// Serialize serializes and item into a bytearray.
func Serialize(item interface{}) []byte {
	return nil
}

// Deserializes an item from a bytearray.
func Deserialize(b []byte) interface{} {
	return nil
}

// Log instucts the VM to log the given message.
func Log(message string) {}

// Notify an event to the VM.
func Notify(arg interface{}) int {
	return 0
}

// GetTime returns the timestamp of the most recent block.
func GetTime() int {
	return 0
}
