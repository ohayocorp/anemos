'use strict';

const assert = require("./assert.js");
const anemos = require("@ohayocorp/anemos");

exports.default = function () {
    const globalVariable = anemos.globalVariable;
    const globalVariablePointer = anemos.globalVariablePointer;
    const globalVariableNamespace = anemos.ns.globalVariable;
    const property = anemos.globalObject.property;
    const pointer = anemos.globalObject.pointer;

    assert.equal(globalVariable, 1);
    assert.equal(globalVariablePointer, 2);
    assert.equal(globalVariableNamespace, 3);
    assert.equal(property, 4);
    assert.equal(pointer, 5);

    anemos.globalVariable = 11;
    anemos.globalVariablePointer = 12;
    anemos.ns.globalVariable = 13;
    anemos.globalObject.property = 14;
    anemos.globalObject.pointer = 15;
    
    assert.equal(anemos.globalVariable, 11);
    assert.equal(anemos.globalVariablePointer, 12);
    assert.equal(anemos.ns.globalVariable, 13);
    assert.equal(anemos.globalObject.property, 14);
    assert.equal(anemos.globalObject.pointer, 15);

    anemos.globalVariablePointer = null;
    anemos.globalObject.pointer = null;
    assert.equal(anemos.globalVariablePointer, null);
    assert.equal(anemos.globalObject.pointer, null);
};