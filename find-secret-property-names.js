const fs = require("node:fs/promises")

/*
 * This script reads a JSON file and prints the names of all properties that have the x-ms-secret key.
 * The x-ms-secret key is used to mark properties that contain sensitive information in Azure APIs.
 *
 * Example usage:
 * node generate.js path/to/file.json
*/

async function generate(filePath) {
  try {
    const data = await fs.readFile(filePath, "utf-8")
    const json = JSON.parse(data)

    secretPropertyNames = findSecretPropertyNames(json, [])

    for (const secretPropertyName of secretPropertyNames) {
      console.log(secretPropertyName)
    }
  } catch (error) {
    console.error("Error parsing", filePath, error)
    process.exitCode = 1
  }
}

function findSecretPropertyNames(object, secretPropertyNames) {
  if (Array.isArray(object)) {
    object.forEach((element, index) => {
      secretPropertyNames = findSecretPropertyNames(element, secretPropertyNames)
    })
  } else if (typeof object === "object") {
    if (object == null) {
      return secretPropertyNames
    }

    Object.keys(object).forEach((key) => {
      if (object[key] === null) {
        return
      }

      if (object[key].hasOwnProperty("x-ms-secret")) {
        secretPropertyNames.push(key)

        return
      }

      secretPropertyNames = findSecretPropertyNames(object[key], secretPropertyNames)
    })
  }

  return secretPropertyNames
}

if (process.argv.length < 3) {
  console.error("Please provide the path to a JSON file")
  process.exitCode = 1
  return
}

generate(process.argv[2])

