const anemos = require('@ohayocorp/anemos');
// Use anemos-apply-package as the package name since we may not have the exact name of the package.
// When package is installed, an alias with the name 'anemos-apply-package' will be created.
// This allows us to use the package without knowing its exact name.
const package = require('anemos-apply-package');

if (typeof package.add !== 'function') {
    throw new Error('Package does not export an add function');
}

// Create a builder to collect manifests. clusterInfo and environmentType are passed from the native code.
const builder = new anemos.Builder(clusterInfo.Version, clusterInfo.Distribution, environmentType);

for (let i = builder.components.length - 1; i >= 0; i--) {
    // Remove all default components from the builder.
    builder.removeComponent(builder.components[i]);
}

// Call the add function from the script with our builder.
// Options is passed from the native code.
package.add(builder, options);

// Add the apply component to the builder. skipConfirmation and namespace are passed from the native code.
builder.apply({
    skipConfirmation: skipConfirmation,
    namespace: namespace,
});
builder.build();