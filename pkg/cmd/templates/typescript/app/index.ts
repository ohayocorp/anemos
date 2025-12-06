import * as anemos from "@ohayocorp/anemos";

const builder = new anemos.Builder();

const name = "PACKAGE_NAME";
const namespace = "default";
const image = "nginx:latest";

builder.addDocument(
    `
    apiVersion: v1
    kind: Pod
    metadata:
      name: ${name}
      namespace: ${namespace}
      labels:
        app: ${name}
    spec:
      containers:
        - name: nginx
          image: ${image}
          ports:
            - containerPort: 80
    `);

builder.build();