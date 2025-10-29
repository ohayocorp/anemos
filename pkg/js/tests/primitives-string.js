'use strict';

const assert = require("./assert.js");
const anemos = require("@ohayocorp/anemos");
const module = require("@ohayocorp/anemos/module");

const globalVariable = anemos.globalVariable;
const globalVariablePointer = anemos.globalVariablePointer;
const globalVariableModule = module.globalVariable;
const property = anemos.globalObject.property;
const pointer = anemos.globalObject.pointer;

assert.equal(globalVariable, "globalVariable");
assert.equal(globalVariablePointer, "globalVariablePointer");
assert.equal(globalVariableModule, "globalVariableModule");
assert.equal(property, "instanceProperty");
assert.equal(pointer, "instancePointer");

anemos.globalVariable = "newGlobalVariable";
anemos.globalVariablePointer = "newGlobalVariablePointer";
module.globalVariable = "newGlobalVariableModule";
anemos.globalObject.property = "newInstanceProperty";
anemos.globalObject.pointer = "newInstancePointer";

assert.equal(anemos.globalVariable, "newGlobalVariable");
assert.equal(anemos.globalVariablePointer, "newGlobalVariablePointer");
assert.equal(module.globalVariable, "newGlobalVariableModule");
assert.equal(anemos.globalObject.property, "newInstanceProperty");
assert.equal(anemos.globalObject.pointer, "newInstancePointer");

anemos.globalVariablePointer = null;
anemos.globalObject.pointer = null;
assert.equal(anemos.globalVariablePointer, null);
assert.equal(anemos.globalObject.pointer, null);

module.globalVariable = "moduleSet";
assert.equal(anemos.module.globalVariable, "moduleSet");