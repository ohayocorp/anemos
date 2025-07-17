const anemos = require("@ohayocorp/anemos");
const { Options } = require("./options");

/**
 * @typedef {Object} Component
 * @property {Options} [options] - Options for the component.
 */
class Component extends anemos.Component {
    /**
     * @param {Options} options 
     */
    constructor(options) {
        super();
        
        this.options = options ?? {};

        this.addAction(anemos.steps.sanitize, this.sanitize);
        this.addAction(anemos.steps.generateResources, this.generateResources);
    }

    /**
     * @param {anemos.BuildContext} context 
     */
    sanitize = (context) => {
        this.options.name ??= "example-package";
        this.options.namespace ??= "default";
        this.options.image ??= "nginx:latest";
    }

    /**
     * @param {anemos.BuildContext} context 
     */
    generateResources = (context) => {
        const pod = new anemos.Document(
            `pod.yaml`,
            `
            apiVersion: "v1"
            kind: "Pod"
            metadata:
              name: ${this.options.name}
              namespace: ${this.options.namespace}
              labels:
                app: ${this.options.name}
            spec:
              containers:
                - name: "nginx"
                  image: "${this.options.image}"
            `);

        const documentGroup = new anemos.DocumentGroup(this.options.name);
        documentGroup.addDocument(pod);

        context.addDocumentGroup(documentGroup);
    }
}

module.exports = Component;