Pod Hunt
====

# What is it?

A game which you run on your OpenShift 4 cluster. Similar to classic Duck Hunt arcade the player would
shoot ducks. Every time the duck is shot the backend would destroy a random pod, a deployment or a statefulset in a
random namespace.

# Installation

* Update `kustomize-route.yaml` with a route to your cluster

* Run kustomize:
```
oc apply -k .
```

* Wait for builds to complete and the app to rollout

* Open two windows - one with a list of all pods in all namespaces, filtering out `Running` - and
one with "Cluster Operators" filtering out those which are `Available`

* Shoot ducks and watch the cluster go crazy

# Limitations

The game would skip several namespaces by default:

* `openshift-cluster-version` - there is nothing which brings CVO back
* `openshift-console`, `openshift-ingress` - its not fun if frontend can't reach backend or can't display current state
* `pod-hunt` - don't shoot yourself in the foot
* `openshift-etcd` - killing a static pod won't do anything, but its not fair if backend reported that pod is killed but nothing in fact happened

# Multiplayer mode

Since frontend is stateless this game can be played by multiple people simultaneously:

* Convert frontend URL in a QR code (https://www.the-qrcode-generator.com/)
* Let people scan it on mobile phones
* Keep logs from backend to file bugs if issues occur
