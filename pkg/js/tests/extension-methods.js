'use strict';

const assert = require("./assert.js");
const anemos = require("@ohayocorp/anemos");

assert.equal(anemos.object.extensionNoParams(), undefined);
assert.equal(anemos.object.returnProperty(), "test");
assert.equal(anemos.object.returnProperty("-1"), "test-1");