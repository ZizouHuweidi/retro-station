# Deploying to an Istio-enabled cluster

This repository provides an [`istio-manifests`](/istio-manifests) directory containing ingress resources (an Istio `Gateway` and `VirtualService`) needed to expose the app frontend running inside a Kubernetes cluster.

You can apply these resources to your cluster in addition to the `kubernetes-manifests`, then use the Istio IngressGateway's external IP to view the app frontend. See the following instructions for Istio steps.

## Steps

1. [Install Istio](https://istio.io/latest/docs/setup/getting-started/) on your cluster.

2. Enable Istio sidecar proxy injection in the `default` Kubernetes namespace.

   ```sh
   kubectl label namespace default istio-injection=enabled
   ```

3. Apply all the manifests in the `/release` directory. This includes the Istio and Kubernetes manifests.

   ```sh
   kubectl apply -f ./release
   ```

4. Run `kubectl get pods` to see pods are in a healthy and ready state.

5. Find the IP address of your Istio gateway Ingress or Service, and visit the
   application frontend in a web browser.

   ```sh
   INGRESS_HOST="$(kubectl -n istio-system get service istio-ingressgateway \
      -o jsonpath='{.status.loadBalancer.ingress[0].ip}')"
   echo "$INGRESS_HOST"
   ```

   ```sh
   curl -v "http://$INGRESS_HOST"
   ```
