import * as anemos from "@ohayocorp/anemos";

// Create an empty mapping.
const emptyMapping = new anemos.Mapping();

// Create a mapping from simple key-value pairs.
const mappingFromKeyValuePairs = new anemos.Mapping({
    "key1": "value1",
    "key2": "value2",
});

// Create a mapping from an object. Values are converted to YAML nodes.
const mappingFromObject = new anemos.Mapping({
    key1: "value1",
    key2: {
        key1: "value1",
        key2: [
            "value1",
            "value2",
        ],
    },
});

// Parse a YAML string as a mapping.
const mappingParsed = anemos.parseMapping(`
    key1: value1
    key2:
      key1: value1
      key2: value2
    `);

// Modify a mapping.
const mapping = new anemos.Mapping();

// Set simple key-value pairs.
mapping.set("key1", "value1");

// Set a mapping as a value.
mapping.set("key2", {
    key1: "value1",
    key2: "value2",
});

// Set a sequence as a value.
mapping.set("key3", [
    "value1",
    "value2",
]);

// Remove a key from a mapping.
mapping.remove("key1");

// Insert a key-value pair at the given index.
mapping.insert(0, "key1", "value1");