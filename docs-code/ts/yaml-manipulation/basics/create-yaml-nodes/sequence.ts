import * as anemos from "@ohayocorp/anemos";

// Create an empty sequence.
const emptySequence = new anemos.Sequence();

// Create a sequence from an array.
const sequenceFromArray = new anemos.Sequence([
    "value1",
    {
        key1: "mapping",
    },
    [
        "another sequence"
    ],
]);

const sequenceParsed = anemos.parseSequence(`
    - key1: value1
    - key2:
        key1: value1
        key2: value2
    `);