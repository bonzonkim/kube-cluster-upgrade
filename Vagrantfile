Vagrant.configure("2") do |config|
  config.vm.box = "bento/ubuntu-22.04"

  config.vm.define "controller" do |vm1|
    vm1.vm.hostname = "controller"
    vm1.vm.network "private_network", ip: "192.168.64.190"
  end

  config.vm.define "worker1" do |vm2|
    vm2.vm.hostname = "worker"
    vm2.vm.network "private_network", ip: "192.168.64.191"
  end

  config.vm.define "worker2" do |vm3|
    vm3.vm.hostname = "worker"
    vm3.vm.network "private_network", ip: "192.168.64.192"
  end
end
