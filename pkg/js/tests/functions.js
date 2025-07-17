'use strict';

const assert = require("./assert.js");
const anemos = require("@ohayocorp/anemos");

exports.default = function () {
    assert.equal(anemos.noParams(), undefined);

    assert.equal(anemos.returnBool(), true);
    assert.equal(anemos.returnBool(false), false);
    assert.equal(anemos.returnBoolPointer(true), true);
    assert.equal(anemos.returnBoolPointer(null), null);

    assert.equal(anemos.returnInt(), 1);
    assert.equal(anemos.returnInt(2), 2);
    assert.equal(anemos.returnIntPointer(3), 3);
    assert.equal(anemos.returnIntPointer(null), null);

    assert.equal(anemos.returnFloat(), 1.2);
    assert.equal(anemos.returnFloat(2), 2);
    assert.equal(anemos.returnFloatPointer(3.4), 3.4);
    assert.equal(anemos.returnFloatPointer(null), null);

    assert.equal(anemos.returnString(), "test");
    assert.equal(anemos.returnString("test2"), "test2");
    assert.equal(anemos.returnStringPointer("test3"), "test3");
    assert.equal(anemos.returnStringPointer(null), null);
};