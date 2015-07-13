#!/bin/bash

# This is an example script that creates a fully functional vitess cluster.
# It performs the following steps:
# 1. Create a container engine cluster
# 2. Create etcd clusters
# 3. Create vtctld clusters
# 4. Forward vtctld port
# 5. Create vttablet clusters
# 6. Perform vtctl initialization:
#      SetKeyspaceShardingInfo, Rebuild Keyspace, Reparent Shard, Apply Schema
# 7. Create vtgate clusters
# 8. Forward vtgate port

# Customizable parameters
GKE_ZONE=${GKE_ZONE:-'us-central1-b'}
GKE_MACHINE_TYPE=${GKE_MACHINE_TYPE:-'n1-standard-1'}
GKE_CLUSTER_NAME=${GKE_CLUSTER_NAME:-'example'}
GKE_SSD_SIZE_GB=${GKE_SSD_SIZE_GB:-0}
GKE_NUM_NODES=${GKE_NUM_NODES:-0}
VTDATAROOT_VOLUME=${VTDATAROOT_VOLUME:-'/ssd'}

# Get region from zone (everything to last dash)
gke_region=`echo $GKE_ZONE | sed "s/-[^-]*$//"`

export KUBECTL='kubectl'
gcloud config set compute/zone $GKE_ZONE
project_id=`gcloud config list project | sed -n 2p | cut -d " " -f 3`

echo "****************************"
echo "*Creating cluster:"
echo "*  Zone: $GKE_ZONE"
echo "*  Machine type: $GKE_MACHINE_TYPE"
echo "*  Num nodes: $GKE_NUM_NODES"
echo "*  SSD Size: $GKE_SSD_SIZE_GB"
echo "*  Cluster name: $GKE_CLUSTER_NAME"
echo "*  Project ID: $project_id"
echo "****************************"
gcloud beta container clusters create $GKE_CLUSTER_NAME --machine-type $GKE_MACHINE_TYPE --num-nodes $GKE_NUM_NODES
gcloud config set container/cluster $GKE_CLUSTER_NAME

if [ $GKE_SSD_SIZE_GB -gt 0 ]
then
  echo Creating SSDs and attaching to container engine nodes
  i=1
  for nodename in `$KUBECTL get nodes --no-headers | awk '{print $1}'`; do
    diskname=$GKE_CLUSTER_NAME-vt-ssd-$i
    gcloud compute disks create $diskname --type=pd-ssd --size=${GKE_SSD_SIZE_GB}GB
    gcloud compute instances attach-disk $nodename --disk $diskname
    gcloud compute ssh $nodename --zone=$GKE_ZONE --command "sudo mkdir /ssd; sudo /usr/share/google/safe_format_and_mount -m \"mkfs.ext4 -F\" /dev/disk/by-id/google-persistent-disk-1 ${VTDATAROOT_VOLUME} &"
    let i=i+1
  done
fi

if [ -n "$NEWRELIC_LICENSE_KEY" -a $GKE_SSD_SIZE_GB -gt 0 ]; then
  i=1
  for nodename in `$KUBECTL get nodes --no-header | awk '{print $1}'`; do
    gcloud compute copy-files newrelic.sh $nodename:~/
    gcloud compute copy-files newrelic_start_agent.sh $nodename:~/
    gcloud compute copy-files newrelic_start_mysql_plugin.sh $nodename:~/
    gcloud compute ssh $nodename --command "bash -c '~/newrelic.sh ${NEWRELIC_LICENSE_KEY}'"
    let i=i+1
  done
fi

echo "****************************"
echo "* Complete!"
echo "****************************"
