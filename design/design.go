package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("error_service", func() {
	Title("Error Service API")
	Description("API for handling error logs and error groups.")
	Server("error_service", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})

var _ = Service("error", func() {
	Description("The error service manages error logs.")

	Method("create", func() {
		Payload(CreateErrorLogFormModel)
		Result(String)
		HTTP(func() {
			POST("/")
			Response(StatusOK)
		})
	})

	Method("updateStatus", func() {
		Payload(LogErrorRequest)
		HTTP(func() {
			POST("/update-status")
			Response(StatusNoContent)
		})
	})

	Method("markResolved", func() {
		Payload(LogErrorRequest)
		HTTP(func() {
			POST("/resolve")
			Response(StatusNoContent)
		})
	})

	Method("markInProgress", func() {
		Payload(LogErrorRequest)
		HTTP(func() {
			POST("/inprogress")
			Response(StatusNoContent)
		})
	})

	Method("markNew", func() {
		Payload(LogErrorRequest)
		HTTP(func() {
			POST("/new")
			Response(StatusNoContent)
		})
	})

	Method("getAll", func() {
		Result(CollectionOf(LogError))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	Method("delete", func() {
		Payload(LogErrorRequest)
		HTTP(func() {
			DELETE("/")
			Response(StatusNoContent)
		})
	})
})

var _ = Service("error_group", func() {
	Description("The error group service manages grouped error logs.")

	Method("list", func() {
		Result(CollectionOf(ErrorGroupDto))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	Method("delete", func() {
		Payload(ErrorGroupRequest)
		HTTP(func() {
			POST("/delete")
			Response(StatusNoContent)
		})
	})
})

var CreateErrorLogFormModel = Type("CreateErrorLogFormModel", func() {
	Attribute("message", String, "Error message", func() { Example("An unexpected error occurred.") })
	Attribute("stacktrace", String, "Stack trace of the error", func() { Example("at line 10...") })
	Required("message", "stacktrace")
})

var LogErrorRequest = Type("LogErrorRequest", func() {
	Attribute("id", String, "Unique error ID", func() { Example("123e4567-e89b-12d3-a456-426614174000") })
	Attribute("status", String, "New status", func() { Example("Resolved") })
	Required("id", "status")
})

var LogError = Type("LogError", func() {
	Attribute("id", String)
	Attribute("message", String)
	Attribute("stacktrace", String)
})

var ErrorGroupDto = Type("ErrorGroupDto", func() {
	Attribute("id", String)
	Attribute("logDetails", String)
	Attribute("logLevel", String)
	Attribute("stacktrace", String)
	Attribute("customer", String)
	Attribute("applicationName", String)
	Attribute("path", String)
	Attribute("firstOccurance", String)
	Attribute("lastOccurance", String)
	Attribute("count", Int)
	Attribute("status", String)
})

var ErrorGroupRequest = Type("ErrorGroupRequest", func() {
	Attribute("id", String, "Unique error group ID", func() { Example("123e4567-e89b-12d3-a456-426614174000") })
	Required("id")
})

