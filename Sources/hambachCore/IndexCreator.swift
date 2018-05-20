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

    func createIndex(content: [Content]) throws -> String
    {
        let article = try self.replacePlaceholder(content: content)

        return article
    }

    private func replacePlaceholder(content: [Content]) throws -> String
    {
        let articleLayoutString = try self.articleLayout.getTemplate()

        return articleLayoutString.replacingOccurrences(of: "%carousel%", with: try self.createCarousel(content: content))
    }

    private func createCarousel(content: [Content])  throws -> String
    {
        var carousel = ""
        var image = "/images/hambach_logo.png"
        let content = Array(content.suffix(3))
        var counter = 0
        var active = ""
        for contentItem in content {
            if (contentItem.titleImage != "") {
                image = contentItem.titleImage
            }
            if (counter == 0) {
                active = "active"
            } else {
                active = ""
            }
            guard let id = contentItem.contentId else {
                throw Abort(.notFound, reason: "No content id found.")
            }
            carousel += "<div class=\"carousel-item " + active + "\" style=\"background-image: url('" + image +  "')\">"
            carousel += "<div class=\"carousel-caption d-block bg-success text-white\"><h3>"
            carousel += "<a href=\"/article/\(id)\" class=\"text-white\">" + contentItem.title + "</a>"
            carousel += "</h3></div></div>"
            counter += 1
        }

        return carousel
    }
}
