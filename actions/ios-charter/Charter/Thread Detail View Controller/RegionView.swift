//
//  RegionView.swift
//  CollapsibleTextView
//
//  Created by Matthew Palmer on 6/02/2016.
//  Copyright © 2016 Matthew Palmer. All rights reserved.
//

import UIKit

protocol RegionViewDataSource: class {
    func numberOfRegionsInRegionView(regionView: RegionView) -> Int
    func regionView(regionView: RegionView, viewForRegionAtIndex: Int) -> UIView
}

protocol RegionViewDelegate: class {
    func regionView(regionView: RegionView, didFinishReplacingRegionAtIndex: Int)
}

/// A vertical stack view abstracted into regions.
class RegionView: UIView {
    weak var dataSource: RegionViewDataSource? {
        didSet {
            if dataSource !== oldValue {
                reloadData()
            }
        }
    }
    
    weak var delegate: RegionViewDelegate?
    
    lazy var stackView: UIStackView = {
        let stackView = UIStackView()
        stackView.backgroundColor = .redColor()
        stackView.axis = .Vertical
        return stackView
    }()
    
    override init(frame: CGRect) {
        super.init(frame: frame)
        postInit()
    }
    
    override class func requiresConstraintBasedLayout() -> Bool { return true }
    
    required init?(coder aDecoder: NSCoder) {
        super.init(frame: CGRectZero)
        postInit()
    }
    
    private func postInit() {
        addSubview(stackView)
        translatesAutoresizingMaskIntoConstraints = false
        stackView.translatesAutoresizingMaskIntoConstraints = false
    }
    
    override func updateConstraints() {
        super.updateConstraints()
        
        let top = NSLayoutConstraint(item: stackView, attribute: .Top, relatedBy: .Equal, toItem: self, attribute: .Top, multiplier: 1.0, constant: 0.0)
        let left = NSLayoutConstraint(item: stackView, attribute: .Left, relatedBy: .Equal, toItem: self, attribute: .Left, multiplier: 1.0, constant: 0.0)
        let right = NSLayoutConstraint(item: stackView, attribute: .Right, relatedBy: .Equal, toItem: self, attribute: .Right, multiplier: 1.0, constant: 0.0)
        let bottom = NSLayoutConstraint(item: stackView, attribute: .Bottom, relatedBy: .Equal, toItem: self, attribute: .Bottom, multiplier: 1.0, constant: 0.0)
        
        addConstraints([top, left, right, bottom])
    }
    
    func reloadData() {
        guard let dataSource = dataSource else { return }
        
        stackView.subviews.forEach { $0.removeFromSuperview() }
        
        let numberOfRegions = dataSource.numberOfRegionsInRegionView(self)
        for index in 0..<numberOfRegions {
            let region = dataSource.regionView(self, viewForRegionAtIndex: index)
            stackView.addArrangedSubview(region)
        }
    }
    
    func replaceRegionAtIndex(index: Int, withView replacementView: UIView) {
        let originalView = stackView.arrangedSubviews[index]
        replacementView.hidden = true
        originalView.hidden = true
        
        self.stackView.insertArrangedSubview(replacementView, atIndex: index)
        self.stackView.removeArrangedSubview(originalView)
        replacementView.hidden = false
        self.delegate?.regionView(self, didFinishReplacingRegionAtIndex: index)
    }
}
