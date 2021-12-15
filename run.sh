#!/bin/bash

# CONTANTS
red=`tput setaf 1`
green=`tput setaf 2`
yellow=`tput setaf 3`
reset=`tput sgr0`

echo "${green}Hello from DevOps!"

echo "${reset}Input an option from the list bellow:"
echo "0 - Install needed homebrew dependencies"
echo "1 - Run localy kubernetes with MiniKube"
echo "2 - Stop localy kubernetes with MiniKube"
echo "3 - List Kubernets PODS"
echo "4 - Build the application image"
echo "5 - run the application image in Docker"
echo "99 - Remove needed homebrew dependencies"

read -r option;

if [ "$option" == "0" ];
then
  echo "${red} we will install some dependencies on your system"
  brew bundle install
fi


if [ "$option" == "1" ];
then
  echo "${red} Let's run the MiniKube Cluster"
  echo "${red} This will set some kubernetes enviroment variables for you, so remember to use the option 2 to gracefully tear down"
  echo -e "${reset}"
  echo "${green} we will build the application using the TAG devops_application:latest"
  docker build -t devops_application:v1 ./src/main
  minikube start
  echo "${green} we will load the image on the minikube image cache to be used"
  minikube cache add devops_application:latest
  eval $(minikube docker-env)
  /bin/bash -c "minikube dashboard" &
  docker build -t devops_application:v1 ./src/main
  docker image list
  helm install redis ./redis
  helm install main ./main
  echo -e "\n ${green}Wait a second for the cluster"
  sleep 1
  echo -e "\n ${green}Wait a second for the cluster - Done"
  echo -e "\n ${green}Listing PODS with: ${red}kubectl get pods --watch - ${green}you can safely use ctrl+c to exit the script now \n${reset}"
  echo -e "${reset}"
  kubectl get pods --watch
fi

if [ "$option" == "2" ];
then
  echo "${green}Let's Delete the MiniKube Cluster"
  echo "${reset}"
  minikube delete
  eval '$(minikube docker-env -u)'
fi

if [ "$option" == "3" ];
then
  echo "${green}Listing kubernets cluster pods"
  echo "${yellow}Tip: you can run: ${red}kubectl get pods ${yellow} - and get the same result"
  echo "${reset}"
  kubectl get pods

fi

if [ "$option" == "4" ];
then
  echo "${green} we will build the application using the TAG devops_application:latest"
  docker build -t devops_application:latest ./src/main
fi

if [ "$option" == "5" ];
then
  echo "${green} we will run the application using the TAG devops_application:latest in Docker"
  docker run devops_application:latest
fi

if [ "$option" == "99" ];
then
  echo "${red} we will UNinstall the dependencies from your system"
  brew bundle install
fi