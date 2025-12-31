package internal

func GetMode(flags []string) []string {
  trimLeading(flags)
  var processedFlags []string

  for _, flag := range flags {
    if !isAllowed(flag) {
      if !contains(processedFlags, "error") {
        processedFlags = append(processedFlags, "error")
      }
    } else {
      if !contains(processedFlags, flag) {
        processedFlags = append(processedFlags, flag)
      }
    }
  }

  return processedFlags
}

// remove leading '-'
func trimLeading(flags []string) {
  for i, flag := range flags {
    if len(flag) > 0 && flag[0] == '-' {
      flags[i] = flag[1:]
    }
  }
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
