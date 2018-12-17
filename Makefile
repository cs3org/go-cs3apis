all:
	git clone https://github.com/cernbox/cs3apis build && cd build && make
	cat prototool_gen_go.yaml >> build/prototool.yaml
	cd build && make generate
	rm -rf build

clean:
	rm -rf build
