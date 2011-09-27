package selenium

var defaultProfile = map[string][string] {
	"app.update.auto": "false",
	"app.update.enabled": "false",
	"browser.startup.page" : "0",
	"browser.download.manager.showWhenStarting": "false",
	"browser.EULA.override": "true",
	"browser.EULA.3.accepted": "true",
	"browser.link.open_external": "2",
	"browser.link.open_newwindow": "2",
	"browser.offline": "false",
	"browser.safebrowsing.enabled": "false",
	"browser.search.update": "false",
	"browser.sessionstore.resume_from_crash": "false",
	"browser.shell.checkDefaultBrowser": "false",
	"browser.tabs.warnOnClose": "false",
	"browser.tabs.warnOnOpen": "false",
	"browser.startup.page": "0",
	"startup.homepage_welcome_url": "\"about:blank\"",
	"devtools.errorconsole.enabled": "true",
	"dom.disable_open_during_load": "false",
	"dom.max_script_run_time": "30",
	"extensions.logging.enabled": "true",
	"extensions.update.enabled": "false",
	"extensions.update.notifyUser": "false",
	"network.manage-offline-status": "false",
	"network.http.max-connections-per-server": "10",
	"network.http.phishy-userpass-length": "255",
	"prompts.tab_modal.enabled": "false",
	"security.fileuri.origin_policy": "3",
	"security.fileuri.strict_origin_policy": "false",
	"security.warn_entering_secure": "false",
	"security.warn_submit_insecure": "false",
	"security.warn_entering_secure.show_once": "false",
	"security.warn_entering_weak": "false",
	"security.warn_entering_weak.show_once": "false",
	"security.warn_leaving_secure": "false",
	"security.warn_leaving_secure.show_once": "false",
	"security.warn_submit_insecure": "false",
	"security.warn_viewing_mixed": "false",
	"security.warn_viewing_mixed.show_once": "false",
	"signon.rememberSignons": "false",
	"toolkit.networkmanager.disable": "true",
	"javascript.options.showInConsole": "true",
	"browser.dom.window.dump.enabled": "true",
	"webdriver_accept_untrusted_certs": "true",
	"webdriver_enable_native_events": "true",
	"dom.max_script_run_time": "30",
}

type FirefoxProfile {
	Root string
}

