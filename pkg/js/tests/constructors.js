'use strict';

const assert = require("./assert.js");
const anemos = require("@ohayocorp/anemos");

exports.default = function () {
    assert.isDefined(new anemos.Test());
    assert.isDefined(new anemos.Test(false, 1, 2.5, "test"));
};