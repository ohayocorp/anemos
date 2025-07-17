'use strict';

const assert = require("./assert.js");
const anemos = require("@ohayocorp/anemos");

exports.default = function () {
    const globalVariable = anemos.globalVariable;
    const globalVariablePointer = anemos.globalVariablePointer;
    const globalVariableNamespace = anemos.ns.globalVariable;
    const property = anemos.globalObject.property;
    const pointer = anemos.globalObject.pointer;

    assert.equal(globalVariable, "globalVariable");
    assert.equal(globalVariablePointer, "globalVariablePointer");
    assert.equal(globalVariableNamespace, "globalVariableNamespace");
    assert.equal(property, "instanceProperty");
    assert.equal(pointer, "instancePointer");
    
    anemos.globalVariable = "newGlobalVariable";
    anemos.globalVariablePointer = "newGlobalVariablePointer";
    anemos.ns.globalVariable = "newGlobalVariableNamespace";
    anemos.globalObject.property = "newInstanceProperty";
    anemos.globalObject.pointer = "newInstancePointer";

    assert.equal(anemos.globalVariable, "newGlobalVariable");
    assert.equal(anemos.globalVariablePointer, "newGlobalVariablePointer");
    assert.equal(anemos.ns.globalVariable, "newGlobalVariableNamespace");
    assert.equal(anemos.globalObject.property, "newInstanceProperty");
    assert.equal(anemos.globalObject.pointer, "newInstancePointer");

    anemos.globalVariablePointer = null;
    anemos.globalObject.pointer = null;
    assert.equal(anemos.globalVariablePointer, null);
    assert.equal(anemos.globalObject.pointer, null);
};