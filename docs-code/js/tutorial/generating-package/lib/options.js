/**
 * @typedef {Object} Options
 * @property {string} [name] - The name of the pod.
 * @property {string} [namespace] - The namespace for the pod.
 * @property {string} [image] - The container image to use.
 */
class Options {
    name;
    namespace;
    image;
}

module.exports = Options;