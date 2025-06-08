Vagrant.configure("2") do |config|
  config.vm.box = "bento/ubuntu-22.04"

  config.vm.define "vm1" do |vm1|
    vm1.vm.hostname = "ubuntu-vm1"
    #vm1.vm.network "private_network", ip: "192.168.56.11"
  end

  config.vm.define "vm2" do |vm2|
    vm2.vm.hostname = "ubuntu-vm2"
    #vm2.vm.network "private_network", ip: "192.168.56.12"
  end
end
