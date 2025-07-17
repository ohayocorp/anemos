'use strict';

const assert = require("./assert.js");
const anemos = require("@ohayocorp/anemos");

exports.default = function () {
    const globalVariable = anemos.globalVariable;
    const globalVariablePointer = anemos.globalVariablePointer;
    const globalVariableNamespace = anemos.ns.globalVariable;
    const property = anemos.globalObject.property;
    const pointer = anemos.globalObject.pointer;

    assert.equal(globalVariable, true);
    assert.equal(globalVariablePointer, true);
    assert.equal(globalVariableNamespace, true);
    assert.equal(property, true);
    assert.equal(pointer, true);
    
    anemos.globalVariable = false;
    anemos.globalVariablePointer = false;
    anemos.ns.globalVariable = false;
    anemos.globalObject.property = false;
    anemos.globalObject.pointer = false;

    assert.equal(anemos.globalVariable, false);
    assert.equal(anemos.globalVariablePointer, false);
    assert.equal(anemos.ns.globalVariable, false);
    assert.equal(anemos.globalObject.property, false);
    assert.equal(anemos.globalObject.pointer, false);
    
    anemos.globalVariablePointer = null;
    anemos.globalObject.pointer = null;
    assert.equal(anemos.globalVariablePointer, null);
    assert.equal(anemos.globalObject.pointer, null);
};