'use strict';

const assert = require("./assert.js");
const anemos = require("@ohayocorp/anemos");

assert.deepEqual(anemos.boolArray, [true, false]);
assert.deepEqual(anemos.intArray, [1, 2]);
assert.deepEqual(anemos.floatArray, [1.2, 2.3]);
assert.deepEqual(anemos.stringArray, ["a", "b"]);
assert.deepEqual(anemos.object.array[0].property, "a");
assert.deepEqual(anemos.object.array[1].property, "b");

anemos.boolArray[0] = false;
anemos.intArray[0] = 2;
anemos.floatArray[0] = 2.3;
anemos.stringArray[0] = "b";
anemos.object.array[0].property = "b";
anemos.object.array[1].property = "a";

assert.deepEqual(anemos.boolArray, [false, false]);
assert.deepEqual(anemos.intArray, [2, 2]);
assert.deepEqual(anemos.floatArray, [2.3, 2.3]);
assert.deepEqual(anemos.stringArray, ["b", "b"]);
assert.deepEqual(anemos.object.array[0].property, "b");
assert.deepEqual(anemos.object.array[1].property, "a");

anemos.boolArray = [false, true];
anemos.intArray = [2, 1];
anemos.floatArray = [2.3, 1.2];
anemos.stringArray = ["b", "a"];
anemos.object.array = [anemos.object.array[1], anemos.object.array[0]];

assert.deepEqual(anemos.boolArray, [false, true]);
assert.deepEqual(anemos.intArray, [2, 1]);
assert.deepEqual(anemos.floatArray, [2.3, 1.2]);
assert.deepEqual(anemos.stringArray, ["b", "a"]);
assert.deepEqual(anemos.object.array[0].property, "a");
assert.deepEqual(anemos.object.array[1].property, "b");

anemos.object.array.pop();
assert.lengthEquals(anemos.object.array, 1);