package app

type Applet interface {
	Start()
	Terminate()
	Join() error
}

func startApplet(applet Applet) {
	applet.Start()
}

func terminateApplet(applet Applet) {
	applet.Terminate()
}
