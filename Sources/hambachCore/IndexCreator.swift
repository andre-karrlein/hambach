import Vapor
import FluentMySQL
import Foundation

public class IndexCreator
{
    var articleLayout: HtmlTemplate

    init(articleLayout: HtmlTemplate)
    {
        self.articleLayout = articleLayout
    }

    func createIndex(content: [Content], type: String) throws -> String
    {
        let article = try self.replacePlaceholder(content: content, type: type)

        return article
    }

    private func replacePlaceholder(content: [Content], type: String) throws -> String
    {
        let articleLayoutString = try self.articleLayout.getTemplate()

        return articleLayoutString.replacingOccurrences(of: "%carousel%", with: try self.createCarousel(content: content, type: type))
    }

    private func createCarousel(content: [Content], type: String)  throws -> String
    {
        var carousel = ""
        var counter = 0
        var active = ""
        var image = ""
        for contentItem in content.reversed() {
            if (contentItem.type != "article" || contentItem.category != type) {
                continue
            }
            if (counter == 4) {
                break
            }
            if (contentItem.titleImage != "") {
                image = contentItem.titleImage
            } else {
                image = "/images/hambach_logo.png"
            }
            if (counter == 0) {
                active = "active"
            } else {
                active = ""
            }
            guard let id = contentItem.contentId else {
                throw Abort(.notFound, reason: "No content id found.")
            }
            carousel += "<div class=\"carousel-item " + active + "\">"
            carousel += "<img src=\"" + image + "\" alt=\"....\">"
            if (contentItem.title != "" && contentItem.title != " ") {
                carousel += "<div class=\"carousel-caption d-block bg-success text-white\"><h3>"
                carousel += "<a href=\"/article/\(id)\" class=\"text-white\">" + contentItem.title + "</a></h3>"
            }
            carousel += "</div></div>"
            counter += 1
        }

        return carousel
    }
}
