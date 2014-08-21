# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "ubuntu/trusty64"

  config.vm.provision :shell do |s|
    s.path = "https://gist.githubusercontent.com/todd/d3f4cbdffadf2a6aa60d/raw/c8bb0dd04562befd289d26f81d8ee19f9220eb28/shell_provisioner.sh"
    s.privileged = false
  end
end
