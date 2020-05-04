/*
	This file is part of solidity.

	solidity is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	solidity is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with solidity.  If not, see <http://www.gnu.org/licenses/>.
*/
/**
 * @author Christian <c@ethdev.com>
 * @date 2014
 * Implementation of the accept functions of AST nodes, included by AST.cpp to not clutter that
 * file with these mechanical implementations.
 */

#pragma once

#include <iostream>
#include <libsolidity/ast/AST.h>
#include <libsolidity/ast/ASTVisitor.h>
using namespace std;
namespace dev
{
namespace solidity
{

void SourceUnit::accept(ASTVisitor& _visitor)
{
	cout << "\t\tSourceUnit&ASTVisitor accept\n";
	if (_visitor.visit(*this))
		listAccept(m_nodes, _visitor);
	_visitor.endVisit(*this);
}

void SourceUnit::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tSourceUnit&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
		listAccept(m_nodes, _visitor);
	_visitor.endVisit(*this);
}

void PragmaDirective::accept(ASTVisitor& _visitor)
{
	cout << "\t\tPragmaDirective&ASTVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void PragmaDirective::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tPragmaDirective&ASTConstVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void ImportDirective::accept(ASTVisitor& _visitor)
{
	cout << "\t\tImportDirective&ASTVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void ImportDirective::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tImportDirective&ASTConstVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void ContractDefinition::accept(ASTVisitor& _visitor)
{
	cout << "\t\tContractDefinition&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		listAccept(m_baseContracts, _visitor);
		listAccept(m_subNodes, _visitor);
	}
	_visitor.endVisit(*this);
}

void ContractDefinition::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tContractDefinition&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		listAccept(m_baseContracts, _visitor);
		listAccept(m_subNodes, _visitor);
	}
	_visitor.endVisit(*this);
}

void InheritanceSpecifier::accept(ASTVisitor& _visitor)
{
	cout << "\t\tInheritanceSpecifier&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_baseName->accept(_visitor);
		if (m_arguments)
			listAccept(*m_arguments, _visitor);
	}
	_visitor.endVisit(*this);
}

void InheritanceSpecifier::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tInheritanceSpecifier&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_baseName->accept(_visitor);
		if (m_arguments)
			listAccept(*m_arguments, _visitor);
	}
	_visitor.endVisit(*this);
}

void EnumDefinition::accept(ASTVisitor& _visitor)
{
	cout << "\t\tEnumDefinition&ASTVisitor accept\n";
	if (_visitor.visit(*this))
		listAccept(m_members, _visitor);
	_visitor.endVisit(*this);
}

void EnumDefinition::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tEnumDefinition&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
		listAccept(m_members, _visitor);
	_visitor.endVisit(*this);
}

void EnumValue::accept(ASTVisitor& _visitor)
{
	cout << "\t\tEnumValue&ASTVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void EnumValue::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tEnumValue&ASTConstVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void UsingForDirective::accept(ASTVisitor& _visitor)
{
	cout << "\t\tUsingForDirective&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_libraryName->accept(_visitor);
		if (m_typeName)
			m_typeName->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void UsingForDirective::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tUsingForDirective&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_libraryName->accept(_visitor);
		if (m_typeName)
			m_typeName->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void StructDefinition::accept(ASTVisitor& _visitor)
{
	cout << "\t\tStructDefinition&ASTVisitor accept\n";
	if (_visitor.visit(*this))
		listAccept(m_members, _visitor);
	_visitor.endVisit(*this);
}

void StructDefinition::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tStructDefinition&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
		listAccept(m_members, _visitor);
	_visitor.endVisit(*this);
}

void ParameterList::accept(ASTVisitor& _visitor)
{
	cout << "\t\tParameterList&ASTVisitor accept\n";
	if (_visitor.visit(*this))
		listAccept(m_parameters, _visitor);
	_visitor.endVisit(*this);
}

void ParameterList::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tParameterList&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
		listAccept(m_parameters, _visitor);
	_visitor.endVisit(*this);
}

void FunctionDefinition::accept(ASTVisitor& _visitor)
{
	cout << "\t\tFunctionDefinition&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_parameters->accept(_visitor);
		if (m_returnParameters)
			m_returnParameters->accept(_visitor);
		listAccept(m_functionModifiers, _visitor);
		if (m_body)
			m_body->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void FunctionDefinition::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tFunctionDefinition&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_parameters->accept(_visitor);
		if (m_returnParameters)
			m_returnParameters->accept(_visitor);
		listAccept(m_functionModifiers, _visitor);
		if (m_body)
			m_body->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void VariableDeclaration::accept(ASTVisitor& _visitor)
{
	cout << "\t\tVariableDeclaration&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		if (m_typeName)
			m_typeName->accept(_visitor);
		if (m_value)
			m_value->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void VariableDeclaration::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tVariableDeclaration&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		if (m_typeName)
			m_typeName->accept(_visitor);
		if (m_value)
			m_value->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void ModifierDefinition::accept(ASTVisitor& _visitor)
{
	cout << "\t\tModifierDefinition&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_parameters->accept(_visitor);
		m_body->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void ModifierDefinition::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tModifierDefinition&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_parameters->accept(_visitor);
		m_body->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void ModifierInvocation::accept(ASTVisitor& _visitor)
{
	cout << "\t\tModifierInvocation&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_modifierName->accept(_visitor);
		if (m_arguments)
			listAccept(*m_arguments, _visitor);
	}
	_visitor.endVisit(*this);
}

void ModifierInvocation::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tModifierInvocation&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_modifierName->accept(_visitor);
		if (m_arguments)
			listAccept(*m_arguments, _visitor);
	}
	_visitor.endVisit(*this);
}

void EventDefinition::accept(ASTVisitor& _visitor)
{
	cout << "\t\tEventDefinition&ASTVisitor accept\n";
	if (_visitor.visit(*this))
		m_parameters->accept(_visitor);
	_visitor.endVisit(*this);
}

void EventDefinition::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tEventDefinition&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
		m_parameters->accept(_visitor);
	_visitor.endVisit(*this);
}

void ElementaryTypeName::accept(ASTVisitor& _visitor)
{
	cout << "\t\tElementaryTypeName&ASTVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void ElementaryTypeName::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tElementaryTypeName&ASTConstVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void UserDefinedTypeName::accept(ASTVisitor& _visitor)
{
	cout << "\t\tUserDefinedTypeName&ASTVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void UserDefinedTypeName::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tUserDefinedTypeName&ASTConstVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void FunctionTypeName::accept(ASTVisitor& _visitor)
{
	cout << "\t\tFunctionTypeName&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_parameterTypes->accept(_visitor);
		m_returnTypes->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void FunctionTypeName::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tFunctionTypeName&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_parameterTypes->accept(_visitor);
		m_returnTypes->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void Mapping::accept(ASTVisitor& _visitor)
{
	cout << "\t\tMapping&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_keyType->accept(_visitor);
		m_valueType->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void Mapping::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tMapping&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_keyType->accept(_visitor);
		m_valueType->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void ArrayTypeName::accept(ASTVisitor& _visitor)
{
	cout << "\t\tArrayTypeName&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_baseType->accept(_visitor);
		if (m_length)
			m_length->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void ArrayTypeName::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tArrayTypeName&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_baseType->accept(_visitor);
		if (m_length)
			m_length->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void InlineAssembly::accept(ASTVisitor& _visitor)
{
	cout << "\t\tInlineAssembly&ASTVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void InlineAssembly::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tInlineAssembly&ASTConstVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void Block::accept(ASTVisitor& _visitor)
{
	cout << "\t\tBlock&ASTVisitor accept\n";
	if (_visitor.visit(*this))
		listAccept(m_statements, _visitor);
	_visitor.endVisit(*this);
}

void Block::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tBlock&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
		listAccept(m_statements, _visitor);
	_visitor.endVisit(*this);
}

void PlaceholderStatement::accept(ASTVisitor& _visitor)
{
	cout << "\t\tPlaceholderStatement&ASTVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void PlaceholderStatement::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tPlaceholderStatement&ASTConstVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void IfStatement::accept(ASTVisitor& _visitor)
{
	cout << "\t\tIfStatement&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_condition->accept(_visitor);
		m_trueBody->accept(_visitor);
		if (m_falseBody)
			m_falseBody->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void IfStatement::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tIfStatement&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_condition->accept(_visitor);
		m_trueBody->accept(_visitor);
		if (m_falseBody)
			m_falseBody->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void WhileStatement::accept(ASTVisitor& _visitor)
{
	cout << "\t\tWhileStatement&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_condition->accept(_visitor);
		m_body->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void WhileStatement::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tWhileStatement&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_condition->accept(_visitor);
		m_body->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void ForStatement::accept(ASTVisitor& _visitor)
{
	cout << "\t\tForStatement&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		if (m_initExpression)
			m_initExpression->accept(_visitor);
		if (m_condExpression)
			m_condExpression->accept(_visitor);
		if (m_loopExpression)
			m_loopExpression->accept(_visitor);
		m_body->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void ForStatement::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tForStatement&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		if (m_initExpression)
			m_initExpression->accept(_visitor);
		if (m_condExpression)
			m_condExpression->accept(_visitor);
		if (m_loopExpression)
			m_loopExpression->accept(_visitor);
		m_body->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void Continue::accept(ASTVisitor& _visitor)
{
	cout << "\t\tContinue&ASTVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void Continue::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tContinue&ASTConstVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void Break::accept(ASTVisitor& _visitor)
{
	cout << "\t\tBreak&ASTVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void Break::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tBreak&ASTConstVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void Return::accept(ASTVisitor& _visitor)
{
	cout << "\t\tReturn&ASTVisitor accept\n";
	if (_visitor.visit(*this))
		if (m_expression)
			m_expression->accept(_visitor);
	_visitor.endVisit(*this);
}

void Return::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tReturn&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
		if (m_expression)
			m_expression->accept(_visitor);
	_visitor.endVisit(*this);
}

void Throw::accept(ASTVisitor& _visitor)
{
	cout << "\t\tThrow&ASTVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void Throw::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tThrow&ASTConstVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void EmitStatement::accept(ASTVisitor& _visitor)
{
	cout << "\t\tEmitStatement&ASTVisitor accept\n";
	if (_visitor.visit(*this))
		m_eventCall->accept(_visitor);
	_visitor.endVisit(*this);
}

void EmitStatement::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tEmitStatement&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
		m_eventCall->accept(_visitor);
	_visitor.endVisit(*this);
}

void ExpressionStatement::accept(ASTVisitor& _visitor)
{
	cout << "\t\tExpressionStatement&ASTVisitor accept\n";
	if (_visitor.visit(*this))
		if (m_expression)
			m_expression->accept(_visitor);
	_visitor.endVisit(*this);
}

void ExpressionStatement::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tExpressionStatement&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
		if (m_expression)
			m_expression->accept(_visitor);
	_visitor.endVisit(*this);
}

void VariableDeclarationStatement::accept(ASTVisitor& _visitor)
{
	cout << "\t\tVariableDeclarationStatement&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		for (ASTPointer<VariableDeclaration> const& var: m_variables)
			if (var)
				var->accept(_visitor);
		if (m_initialValue)
			m_initialValue->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void VariableDeclarationStatement::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tVariableDeclarationStatement&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		for (ASTPointer<VariableDeclaration> const& var: m_variables)
			if (var)
				var->accept(_visitor);
		if (m_initialValue)
			m_initialValue->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void Conditional::accept(ASTVisitor& _visitor)
{
	cout << "\t\tConditional&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_condition->accept(_visitor);
		m_trueExpression->accept(_visitor);
		m_falseExpression->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void Conditional::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tConditional&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_condition->accept(_visitor);
		m_trueExpression->accept(_visitor);
		m_falseExpression->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void Assignment::accept(ASTVisitor& _visitor)
{
	cout << "\t\tAssignment&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_leftHandSide->accept(_visitor);
		m_rightHandSide->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void Assignment::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tAssignment&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_leftHandSide->accept(_visitor);
		m_rightHandSide->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void TupleExpression::accept(ASTVisitor& _visitor)
{
	cout << "\t\tTupleExpression&ASTVisitor accept\n";
	if (_visitor.visit(*this))
		for (auto const& component: m_components)
			if (component)
				component->accept(_visitor);
	_visitor.endVisit(*this);
}

void TupleExpression::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tTupleExpression&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
		for (auto const& component: m_components)
			if (component)
				component->accept(_visitor);
	_visitor.endVisit(*this);
}

void UnaryOperation::accept(ASTVisitor& _visitor)
{
	cout << "\t\tUnaryOperation&ASTVisitor accept\n";
	if (_visitor.visit(*this))
		m_subExpression->accept(_visitor);
	_visitor.endVisit(*this);
}

void UnaryOperation::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tUnaryOperation&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
		m_subExpression->accept(_visitor);
	_visitor.endVisit(*this);
}

void BinaryOperation::accept(ASTVisitor& _visitor)
{
	cout << "\t\tBinaryOperation&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_left->accept(_visitor);
		m_right->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void BinaryOperation::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tBinaryOperation&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_left->accept(_visitor);
		m_right->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void FunctionCall::accept(ASTVisitor& _visitor)
{
	cout << "\t\tFunctionCall&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_expression->accept(_visitor);
		listAccept(m_arguments, _visitor);
	}
	_visitor.endVisit(*this);
}

void FunctionCall::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tFunctionCall&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_expression->accept(_visitor);
		listAccept(m_arguments, _visitor);
	}
	_visitor.endVisit(*this);
}

void NewExpression::accept(ASTVisitor& _visitor)
{
	cout << "\t\tNewExpression&ASTVisitor accept\n";
	if (_visitor.visit(*this))
		m_typeName->accept(_visitor);
	_visitor.endVisit(*this);
}

void NewExpression::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tNewExpression&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
		m_typeName->accept(_visitor);
	_visitor.endVisit(*this);
}

void MemberAccess::accept(ASTVisitor& _visitor)
{
	cout << "\t\tMemberAccess&ASTVisitor accept\n";
	if (_visitor.visit(*this))
		m_expression->accept(_visitor);
	_visitor.endVisit(*this);
}

void MemberAccess::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tMemberAccess&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
		m_expression->accept(_visitor);
	_visitor.endVisit(*this);
}

void IndexAccess::accept(ASTVisitor& _visitor)
{
	cout << "\t\tIndexAccess&ASTVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_base->accept(_visitor);
		if (m_index)
			m_index->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void IndexAccess::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tIndexAccess&ASTConstVisitor accept\n";
	if (_visitor.visit(*this))
	{
		m_base->accept(_visitor);
		if (m_index)
			m_index->accept(_visitor);
	}
	_visitor.endVisit(*this);
}

void Identifier::accept(ASTVisitor& _visitor)
{
	cout << "\t\tIdentifier&ASTVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void Identifier::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tIdentifier&ASTConstVisitor accept\n";
	cout << "\t\t\tname is "<<*m_name<<"\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void ElementaryTypeNameExpression::accept(ASTVisitor& _visitor)
{
	cout << "\t\tElementaryTypeNameExpression&ASTVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void ElementaryTypeNameExpression::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tElementaryTypeNameExpression&ASTConstVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void Literal::accept(ASTVisitor& _visitor)
{
	cout << "\t\tLiteral&ASTVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

void Literal::accept(ASTConstVisitor& _visitor) const
{
	cout << "\t\tLiteral&ASTConstVisitor accept\n";
	_visitor.visit(*this);
	_visitor.endVisit(*this);
}

}
}
