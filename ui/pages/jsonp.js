(function () {
	const response = {
		data: {
			type: "page",
			title: "标题",
			body: "javascripte page..."
		},
		status: 0
	}

	window.jsonpCallback && window.jsonpCallback(response);
})();
