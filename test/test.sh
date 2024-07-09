#!/bin/bash

# Create the temporary directory if it doesn't exist
mkdir -p ./test/tmp

# Function to compare the output of system xxd and custom ./bin/xxd
compare_xxd() {
  local file_to_test="$1"
  shift
  local xxd_options=($@)

  # Check if the file exists
  if [ ! -f "$file_to_test" ]; then
    echo "File $file_to_test does not exist."
    return 1
  fi

  # Check if both commands are available
  if ! command -v xxd &> /dev/null; then
    echo "System xxd command not found."
    return 1
  fi

  if [ ! -x "./bin/xxd" ]; then
    echo "Custom ./bin/xxd command not found or is not executable."
    return 1
  fi

  # Run both commands and save the output to temporary files
  xxd $xxd_options $file_to_test > ./test/tmp/system_xxd_output.txt
  ./bin/xxd $xxd_options $file_to_test > ./test/tmp/custom_xxd_output.txt

  # Compare the outputs
  diff ./test/tmp/system_xxd_output.txt ./test/tmp/custom_xxd_output.txt > ./test/tmp/xxd_diff.txt

  if [ -s ./test/tmp/xxd_diff.txt ]; then
    echo "Differences found for file $file_to_test with options '$xxd_options':"
    cat ./test/tmp/xxd_diff.txt
  else
    echo "No differences found for file $file_to_test with options '$xxd_options'."
  fi

  # Clean up temporary files
  rm ./test/tmp/system_xxd_output.txt ./test/tmp/custom_xxd_output.txt ./test/tmp/xxd_diff.txt
}

# Example usage
compare_xxd "./test/data/very_small.txt"
compare_xxd "./test/data/very_small.txt" -c 8
compare_xxd "./test/data/very_small.txt" -c 12 -g 8
compare_xxd "./test/data/very_small.txt" -l 20
compare_xxd "./test/data/very_small.txt" -l 20 -c 10


compare_xxd "./test/data/small.txt"
compare_xxd "./test/data/small.txt" -c 8
compare_xxd "./test/data/small.txt" -c 12 -g 8
compare_xxd "./test/data/small.txt" -l 20
compare_xxd "./test/data/small.txt" -l 20 -c 10


compare_xxd "./test/data/medium.txt"
compare_xxd "./test/data/medium.txt" -c 8
compare_xxd "./test/data/medium.txt" -c 12 -g 8
compare_xxd "./test/data/medium.txt" -l 20
compare_xxd "./test/data/medium.txt" -l 20 -c 10


compare_xxd "./test/data/large.bin"
compare_xxd "./test/data/large.bin" -c 8
compare_xxd "./test/data/large.bin" -c 12 -g 8
compare_xxd "./test/data/large.bin" -l 20
compare_xxd "./test/data/large.bin" -l 20 -c 10

rmdir ./test/tmp 

