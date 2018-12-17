package main

import (
	"fmt"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var mainwin *ui.Window
var loginform *ui.Window
var registerform *ui.Window

func makeBasicControlsPage() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)

	hbox.Append(ui.NewButton("Button"), false)
	hbox.Append(ui.NewCheckbox("Checkbox"), false)

	vbox.Append(ui.NewLabel("This is a label. Right now, labels can only span one line."), false)

	vbox.Append(ui.NewHorizontalSeparator(), false)

	group := ui.NewGroup("Entries")
	group.SetMargined(true)
	vbox.Append(group, true)

	group.SetChild(ui.NewNonWrappingMultilineEntry())

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	group.SetChild(entryForm)

	entryForm.Append("Entry", ui.NewEntry(), false)
	entryForm.Append("Password Entry", ui.NewPasswordEntry(), false)
	entryForm.Append("Search Entry", ui.NewSearchEntry(), false)
	entryForm.Append("Multiline Entry", ui.NewMultilineEntry(), true)
	entryForm.Append("Multiline Entry No Wrap", ui.NewNonWrappingMultilineEntry(), true)

	return vbox
}

func makeNumbersPage() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	group := ui.NewGroup("Numbers")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	spinbox := ui.NewSpinbox(0, 100)
	slider := ui.NewSlider(0, 100)
	pbar := ui.NewProgressBar()
	spinbox.OnChanged(func(*ui.Spinbox) {
		slider.SetValue(spinbox.Value())
		pbar.SetValue(spinbox.Value())
	})
	slider.OnChanged(func(*ui.Slider) {
		spinbox.SetValue(slider.Value())
		pbar.SetValue(slider.Value())
	})
	vbox.Append(spinbox, false)
	vbox.Append(slider, false)
	vbox.Append(pbar, false)

	ip := ui.NewProgressBar()
	ip.SetValue(-1)
	vbox.Append(ip, false)

	group = ui.NewGroup("Lists")
	group.SetMargined(true)
	hbox.Append(group, true)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	group.SetChild(vbox)

	cbox := ui.NewCombobox()
	cbox.Append("Combobox Item 1")
	cbox.Append("Combobox Item 2")
	cbox.Append("Combobox Item 3")
	vbox.Append(cbox, false)

	ecbox := ui.NewEditableCombobox()
	ecbox.Append("Editable Item 1")
	ecbox.Append("Editable Item 2")
	ecbox.Append("Editable Item 3")
	vbox.Append(ecbox, false)

	rb := ui.NewRadioButtons()
	rb.Append("Radio Button 1")
	rb.Append("Radio Button 2")
	rb.Append("Radio Button 3")
	vbox.Append(rb, false)

	return hbox
}

func makeDataChoosersPage() ui.Control {
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, false)

	vbox.Append(ui.NewDatePicker(), false)
	vbox.Append(ui.NewTimePicker(), false)
	vbox.Append(ui.NewDateTimePicker(), false)
	vbox.Append(ui.NewFontButton(), false)
	vbox.Append(ui.NewColorButton(), false)

	hbox.Append(ui.NewVerticalSeparator(), false)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	grid := ui.NewGrid()
	grid.SetPadded(true)
	vbox.Append(grid, false)

	button := ui.NewButton("Open File")
	entry := ui.NewEntry()
	entry.SetReadOnly(true)
	button.OnClicked(func(*ui.Button) {
		filename := ui.OpenFile(mainwin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry.SetText(filename)
	})
	grid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry,
		1, 0, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	button = ui.NewButton("Save File")
	entry2 := ui.NewEntry()
	entry2.SetReadOnly(true)
	button.OnClicked(func(*ui.Button) {
		filename := ui.SaveFile(mainwin)
		if filename == "" {
			filename = "(cancelled)"
		}
		entry2.SetText(filename)
	})
	grid.Append(button,
		0, 1, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry2,
		1, 1, 1, 1,
		true, ui.AlignFill, false, ui.AlignFill)

	msggrid := ui.NewGrid()
	msggrid.SetPadded(true)
	grid.Append(msggrid,
		0, 2, 2, 1,
		false, ui.AlignCenter, false, ui.AlignStart)

	button = ui.NewButton("Message Box")
	button.OnClicked(func(*ui.Button) {
		ui.MsgBox(mainwin,
			"This is a normal message box.",
			"More detailed information can be shown here.")
	})
	msggrid.Append(button,
		0, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)
	button = ui.NewButton("Error Box")
	button.OnClicked(func(*ui.Button) {
		ui.MsgBoxError(mainwin,
			"This message box describes an error.",
			"More detailed information can be shown here.")
	})
	msggrid.Append(button,
		1, 0, 1, 1,
		false, ui.AlignFill, false, ui.AlignFill)

	return hbox
}

func mainview() {
	mainwin = ui.NewWindow("libui Control Gallery", 640, 480, true)

	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	tab := ui.NewTab()
	mainwin.SetChild(tab)
	mainwin.SetMargined(true)

	tab.Append("Basic Controls", makeBasicControlsPage())
	tab.SetMargined(0, true)

	tab.Append("Numbers and Lists", makeNumbersPage())
	tab.SetMargined(1, true)

	tab.Append("Data Choosers", makeDataChoosersPage())
	tab.SetMargined(2, true)
	mainwin.Show()
}

func register() {
	registerform = ui.NewWindow("注册", 240, 320, true)
	registerform.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	registerform.SetChild(vbox)

	group := ui.NewGroup("")
	group.SetMargined(true)
	vbox.Append(group, true)

	group.SetChild(ui.NewNonWrappingMultilineEntry())

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	group.SetChild(entryForm)
	user := ui.NewEntry()
	entryForm.Append("账户邮箱", user, false)
	pwd := ui.NewPasswordEntry()

	entryForm.Append("账户密码", pwd, false)
	repwd := ui.NewPasswordEntry()

	entryForm.Append("重复密码", repwd, false)
	vbox.Append(ui.NewHorizontalSeparator(), false)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)
	hbox.Append(ui.NewVerticalSeparator(), false)
	loginBtn := ui.NewButton("登陆")
	hbox.Append(loginBtn, false)
	registerBtn := ui.NewButton("注册")

	hbox.Append(registerBtn, false)

	registerBtn.OnClicked(func(*ui.Button) {
		if user.Text() == "" || pwd.Text() != repwd.Text() {
			ui.MsgBoxError(registerform,
				"注册信息不符合要求",
				"用户信息不能为空且两次密码应该相同")
		} else {
			fmt.Println(user.Text(), pwd.Text(), repwd.Text())
			registerform.Destroy()
			login()
		}

	})

	loginBtn.OnClicked(func(*ui.Button) {
		registerform.Destroy()
		login()

	})

	// vbox.Append(ui.NewLabel("This is a label. Right now, labels can only span one line."), false)

	registerform.Show()

}

func login() {
	loginform = ui.NewWindow("登录", 240, 320, true)
	loginform.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	loginform.SetChild(vbox)

	group := ui.NewGroup("")
	group.SetMargined(true)
	vbox.Append(group, true)

	group.SetChild(ui.NewNonWrappingMultilineEntry())

	entryForm := ui.NewForm()
	entryForm.SetPadded(true)
	group.SetChild(entryForm)
	user := ui.NewEntry()
	entryForm.Append("账户邮箱", user, false)
	pwd := ui.NewPasswordEntry()
	entryForm.Append("账户密码", pwd, false)
	vbox.Append(ui.NewHorizontalSeparator(), false)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	vbox.Append(hbox, false)
	registerBtn := ui.NewButton("注册")
	hbox.Append(registerBtn, false)
	hbox.Append(ui.NewVerticalSeparator(), false)
	loginBtn := ui.NewButton("登录")
	hbox.Append(loginBtn, false)
	hbox.Append(ui.NewCheckbox("记住我"), false)

	registerBtn.OnClicked(func(*ui.Button) {
		loginform.Destroy()
		register()
	})
	user.OnChanged(func (*ui.Entry){
		pwd.SetText("xxxxxx")
	})
	loginBtn.OnClicked(func(*ui.Button) {
		if user.Text() != "tacey" && pwd.Text() != "wong" {
			ui.MsgBoxError(loginform,
				"认证错误",
				"密码错误或用户名和密码不匹配")
		} else {
			loginform.Destroy()
			mainview()
			drawtext()
			table()
			hist()
		}
	})
	// vbox.Append(ui.NewLabel("This is a label. Right now, labels can only span one line."), false)

	loginform.Show()


}

func start() {
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		loginform.Destroy()
		registerform.Destroy()
		return true
	})
	ui.Main(login)
}
