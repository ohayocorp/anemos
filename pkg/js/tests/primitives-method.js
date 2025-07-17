'use strict';

const assert = require("./assert.js");
const anemos = require("@ohayocorp/anemos");

exports.default = function () {
    assert.equal(anemos.test.noParams(), undefined);

    assert.equal(anemos.test.returnBool(), true);
    assert.equal(anemos.test.returnBool(false), false);
    assert.equal(anemos.test.returnBoolPointer(true), true);
    assert.equal(anemos.test.returnBoolPointer(null), null);

    assert.equal(anemos.test.returnInt(), 1);
    assert.equal(anemos.test.returnInt(2), 2);
    assert.equal(anemos.test.returnIntPointer(3), 3);
    assert.equal(anemos.test.returnIntPointer(null), null);

    assert.equal(anemos.test.returnFloat(), 1.2);
    assert.equal(anemos.test.returnFloat(2), 2);
    assert.equal(anemos.test.returnFloatPointer(3.4), 3.4);
    assert.equal(anemos.test.returnFloatPointer(null), null);

    assert.equal(anemos.test.returnString(), "test");
    assert.equal(anemos.test.returnString("test2"), "test2");
    assert.equal(anemos.test.returnStringPointer("test3"), "test3");
    assert.equal(anemos.test.returnStringPointer(null), null);
};