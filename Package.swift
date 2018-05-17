// swift-tools-version:4.0
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "hambach",
    dependencies: [
        .package(url: "https://github.com/vapor/vapor.git", from: "3.0.0"),
        .package(url: "https://github.com/vapor/fluent-mysql.git", from: "3.0.0-rc"),
    ],
    targets: [
        .target(
            name: "hambach",
            dependencies: ["hambachCore"]
        ),
        .target(
            name: "hambachCore",
            dependencies: ["Vapor", "FluentMySQL"]
        ),
        .testTarget(
            name: "hambachCoreTests",
            dependencies: ["hambachCore"]
        )
    ]
)
