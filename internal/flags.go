package internal

func GetMode(flags []string) []string {
  var processedFlags []string

  for _, flag := range flags {
    if !isAllowed(flag) {
      if !contains(processedFlags, "error") {
        processedFlags = append(processedFlags, "error")
      }
    } else {
      processedFlags = append(processedFlags, flag)
    }
  }

  return processedFlags
}

// check if flag is allowed
func isAllowed(flag string) bool {
  allowedFlags := []string{"help", "version", "add", "list", "remove"}
  return contains(allowedFlags, flag)
}

func contains(arr []string, flag string) bool {
  for _, elem := range arr {
    if elem == flag {
      return true
    }
  }
  return false
}
