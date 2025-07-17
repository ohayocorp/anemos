'use strict';

const assert = require("./assert.js");
const anemos = require("@ohayocorp/anemos");

exports.default = function () {
    assert.equal(anemos.object.bool, true);
    assert.equal(anemos.object.int, 1);
    assert.equal(anemos.object.float, 1.2);
    assert.equal(anemos.object.string, "a");
    assert.equal(anemos.object.array[0].property, "a");
    assert.equal(anemos.object.array[1].property, "b");

    const test = anemos.mapParam(anemos.object);

    assert.equal(test.bool, true);
    assert.equal(test.int, 1);
    assert.equal(test.float, 1.2);
    assert.equal(test.string, "a");
    assert.equal(test.array[0].property, "a");
    assert.equal(test.array[1].property, "b");

    anemos.object = { bool: false, int: 2, float: 2.3, string: "b", array: [{ property: "c" }, { property: "d" }] };

    assert.equal(anemos.object.bool, false);
    assert.equal(anemos.object.int, 2);
    assert.equal(anemos.object.float, 2.3);
    assert.equal(anemos.object.string, "b");
    assert.equal(anemos.object.array[0].property, "c");
    assert.equal(anemos.object.array[1].property, "d");
};