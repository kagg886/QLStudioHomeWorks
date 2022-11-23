function loadFriendList() {
	let nameList = document.getElementById("nameList");
	for (let i = 0; i < localStorage.length; i++) {
		let unit = document.createElement("a");
		let name = localStorage.key(i);
		unit.setAttribute("href", "./myinfo.html?i=" + name);
		unit.innerText = name;
		nameList.appendChild(unit);
	}
}

function loadFriendInfo() {
	let name = decodeURI(window.location.href).split("?i=")[1];
	document.title = name + "的介绍";
	document.getElementById("edit_name").innerText = localStorage.getItem(name);
}

function submit() {
	let name = document.getElementById("edit_name").value;
	let value = document.getElementById("edit_value").value;
	if (name == "" || value == "") {
		window.alert("名字/值不能为空");
		return;
	}
	localStorage.setItem(name,value);
	if (localStorage.getItem(name) != "") {
		window.alert("修改成功");
		return;
	}
	window.alert("设置成功");
}