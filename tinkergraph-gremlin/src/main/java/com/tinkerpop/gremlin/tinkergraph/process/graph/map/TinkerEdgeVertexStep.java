package com.tinkerpop.gremlin.tinkergraph.process.graph.map;

import com.tinkerpop.gremlin.process.Traversal;
import com.tinkerpop.gremlin.process.graph.map.EdgeVertexStep;
import com.tinkerpop.gremlin.structure.Direction;
import com.tinkerpop.gremlin.tinkergraph.structure.TinkerEdge;
import com.tinkerpop.gremlin.tinkergraph.structure.TinkerHelper;

import java.util.Iterator;

/**
 * @author Marko A. Rodriguez (http://markorodriguez.com)
 */
public class TinkerEdgeVertexStep extends EdgeVertexStep {

    public TinkerEdgeVertexStep(final Traversal traversal, final Direction direction) {
        super(traversal, direction);
        this.setFunction(holder -> (Iterator) TinkerHelper.getVertices(((TinkerEdge) holder.get()), direction));
    }
}