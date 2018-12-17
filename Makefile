all:
	git clone https://github.com/cernbox/cs3apis build && cd build && make
	cat prototool_gen_go.yaml >> build/prototool.yaml
	cd build && prototool generate
	rm -rf build
	# change import paths
	egrep -lR 'github.com/cernbox/go-cs3apis/cs3/' | xargs sed -i -e 's|github.com/cernbox/go-cs3apis/cs3/|github.com/cernbox/go-cs3apis/cs3/|g'

clean:
	rm -rf build
