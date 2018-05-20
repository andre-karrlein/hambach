import Vapor

let factory = Factory()

public func routes(_ router: Router) throws
{
    router.get("/") { request -> Future<HTTPResponse> in
        return request.withPooledConnection(to: .mysql) { db -> Future<HTTPResponse> in
            let responseCreator = factory.createResponseCreator(template: "Index")
            return db.query(Content.self).all().map(to: HTTPResponse.self) { content in
                return try responseCreator.createResponse(content: content, page: "Index")
            }
        }
    }

    router.get("sportheim") { request -> Future<HTTPResponse> in
        return request.withPooledConnection(to: .mysql) { db -> Future<HTTPResponse> in
            let responseCreator = factory.createResponseCreator(template: "Carousel")
            return try Content.find(1, on: db).map(to: HTTPResponse.self) { content in
                guard let content = content else {
                    throw Abort(.notFound, reason: "Could not find content.")
                }
                return try responseCreator.createResponse(content: [content], page: "Article")
            }
        }
    }

    router.get("article", Int.parameter) { request -> Future<HTTPResponse> in
        let id = try request.parameters.next(Int.self)

        return request.withPooledConnection(to: .mysql) { db -> Future<HTTPResponse> in
            return try Content.find(id, on: db).map(to: HTTPResponse.self) { content in
                guard let content = content else {
                    throw Abort(.notFound, reason: "Could not find content.")
                }
                let responseCreator = factory.createResponseCreator(template: content.type.capitalized)
                return try responseCreator.createResponse(content: [content], page: "Article")
            }
        }
    }
}
