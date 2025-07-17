const anemos = require("@ohayocorp/anemos");
const Component = require("./component");
const Options = require("./options");

/**
 * @param {anemos.Builder} builder 
 * @param {Options} options 
 * @returns {anemos.Component}
 */
function add(builder, options) {
    const component = new Component(options);
    builder.addComponent(component);

    return component;
}

module.exports = {
    add,
    Options
};